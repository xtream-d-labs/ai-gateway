package auth

import (
	"errors"
	"strings"

	"github.com/SermoDigital/jose/jws"
	"github.com/go-openapi/swag"
)

// Principal principal
type Principal struct {
	Username string
}

var invalidToken error

func init() {
	invalidToken = errors.New("Invalid token")
}

// RequestToPrincipal returns principal from token string in the request 'Authorization' header
func RequestToPrincipal(token string) (*Principal, error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	if swag.IsZero(token) {
		return nil, invalidToken
	}
	jwt, err := jws.ParseJWT([]byte(token))
	if jwt == nil || err != nil {
		return nil, invalidToken
	}
	err = jwt.Validate(rsaPublic, signingMethod)
	if err != nil {
		switch err {
		case jws.ErrNoTokenInRequest:
			return nil, invalidToken
		case jws.ErrIsNotJWT:
			return nil, invalidToken
		}
		if strings.Contains(err.Error(), "crypto/rsa") {
			return nil, invalidToken
		}
	}
	data := map[string]interface{}(jwt.Claims())
	if value, claimsFound := data[sessionKey]; claimsFound {
		if session, ok := value.(map[string]interface{}); ok {
			return &Principal{
				Username: findString(session, "docker_username"),
			}, nil
		}
	}
	return nil, invalidToken
}
