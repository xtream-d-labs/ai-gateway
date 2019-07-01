package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/dgraph-io/badger"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/log"
)

// Session defines JWT payload
type Session struct {
	DockerUsername string `json:"docker_username,omitempty"`
	DockerPassword string `json:"-"`
	NgcEmail       string `json:"ngc_email,omitempty"`
	NgcPassword    string `json:"-"`
	NgcApikey      string `json:"-"`
	NgcSession     string `json:"-"`
	K8sConfig      string `json:"-"`
	RescaleKey     string `json:"-"`
}

// Credentials stores third-party credentials
type Credentials struct {
	Base       *models.Configuration `json:"base,omitempty"`
	NgcSession string                `json:"ngc_session,omitempty"`
}

const (
	sessionKey = "claims"
	anonymous  = "anonymous"
)

// ToJWT translate session to JWT string
func (s *Session) ToJWT() (string, error) {
	expiration := 0 // to be filled @ jwt.go L49
	jwt, e := toJWT(s.DockerUsername, expiration, &map[string]interface{}{sessionKey: *s})
	if e != nil {
		log.Error("make-jwt", e, nil)
	}
	return jwt, e
}

// RetrieveSession retrieves the session itself from a HTTP request
func RetrieveSession(req *http.Request) (*Session, error) {
	return retrieveOptionalSession(req, true)
}

// retrieveOptionalSession retrieves the session itself from a HTTP request
func retrieveOptionalSession(req *http.Request, checkValidationError bool) (*Session, error) {
	data, jwtFound, err := retrieveDataFromJWT(req)

	// JWT does not exist in the HTTP request
	if data == nil || !jwtFound {
		return nil, nil
	}
	// JWT does exist but it's invalid
	if err != nil {
		if checkValidationError {
			return nil, err
		}
		// Even though 'checkValidationError' is false, following errors will not be ignored
		switch err {
		case jws.ErrCannotValidate, jws.ErrMismatchedAlgorithms, crypto.ErrInvalidKey:
			return nil, err
		}
		if strings.Contains(err.Error(), "crypto/rsa") {
			return nil, err
		}
	}
	if value, claimsFound := data[sessionKey]; claimsFound {
		if session, ok := value.(map[string]interface{}); ok {
			creds := FindCredentials(findString(session, "docker_username"))
			return creds.ToSession(), nil
		}
	}
	return nil, fmt.Errorf("claims has been broken")
}

func findString(data map[string]interface{}, key string) string {
	if value, found := data[key]; found {
		return toString(value)
	}
	return ""
}

func toString(value interface{}) string {
	if candidate, ok := value.(string); ok {
		return candidate
	}
	return ""
}

// FindCredentials returns creds from local DB
func FindCredentials(username string) *Credentials {
	creds := &Credentials{
		Base: &models.Configuration{
			DockerUsername: username,
		},
	}
	if username == "" {
		username = anonymous
	}
	db.GetValue(func(txn *badger.Txn) error {
		value, err := txn.Get([]byte(username))
		if err != nil {
			return nil
		}
		marshaled, err := value.Value()
		if err != nil {
			return err
		}
		if err := json.Unmarshal(marshaled, &creds); err != nil {
			return err
		}
		return nil
	})
	return creds
}

// Save to local DB
func (c *Credentials) Save() error {
	bytes, err := json.Marshal(c)
	if err != nil {
		return err
	}
	username := c.Base.DockerUsername
	if username == "" {
		username = anonymous
	}
	return db.SetValue(func(txn *badger.Txn) error {
		return txn.Set([]byte(username), bytes)
	})
}

// ToSession translate credentials to session
func (c *Credentials) ToSession() *Session {
	username := c.Base.DockerUsername
	if username == anonymous {
		username = ""
	}
	return &Session{
		DockerUsername: username,
		DockerPassword: c.Base.DockerPassword,
		NgcEmail:       c.Base.NgcEmail.String(),
		NgcPassword:    c.Base.NgcPassword,
		NgcApikey:      c.Base.NgcApikey,
		NgcSession:     c.NgcSession,
		K8sConfig:      c.Base.K8sConfig,
		RescaleKey:     c.Base.RescaleKey,
	}
}
