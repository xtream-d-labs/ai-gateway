package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/rescale-labs/scaleshift/api/src/config"
)

var (
	rsaPrivate    *rsa.PrivateKey
	rsaPublic     *rsa.PublicKey
	signingMethod = crypto.SigningMethodRS256
)

func init() {
	if bytes, err := ioutil.ReadFile(config.Config.JwsPrivateKey); err == nil {
		rsaPrivate, _ = crypto.ParseRSAPrivateKeyFromPEM(bytes)
	}
	if bytes, err := ioutil.ReadFile(config.Config.JwsPublicKey); err == nil {
		rsaPublic, _ = crypto.ParseRSAPublicKeyFromPEM(bytes)
	}
}

func toJWT(userID string, expiration int, content *map[string]interface{}) (string, error) {
	if err := checkRSAKeys(); err != nil {
		return "", err
	}
	return serializeAsJWT(jwsClaims(
		config.Config.JwtIssuer,
		config.Config.JwtAudience,
		userID,
		expiration,
		content,
	))
}

func jwsClaims(issuer, audience, userID string, expiration int, content *map[string]interface{}) jws.Claims {
	claims := jws.Claims{}
	claims.Set("alg", signingMethod.Name)
	claims.SetIssuer(issuer)
	claims.SetSubject(userID)
	claims.SetAudience(audience)
	now := time.Now().UTC()
	claims.SetIssuedAt(now)
	if expiration < 1 {
		expiration = config.Config.JwtExpiration
	}
	claims.SetExpiration(now.Add(time.Duration(expiration) * time.Second))
	claims.Set("uid", userID)
	if content != nil {
		for key, value := range *content {
			claims.Set(key, value)
		}
	}
	return claims
}

func serializeAsJWT(claims jws.Claims) (string, error) {
	jwt, err := jws.NewJWT(claims, signingMethod).Serialize(rsaPrivate)
	if err == nil {
		return fmt.Sprintf("Bearer %s", jwt), nil
	}
	return "", err
}

func retrieveDataFromJWT(req *http.Request) (map[string]interface{}, bool, error) {
	if err := checkRSAKeys(); err != nil {
		return nil, false, err
	}
	jwt, jwtFound, err := jwtValidate(req)
	if jwt == nil {
		return nil, jwtFound, err
	}
	return map[string]interface{}(jwt.Claims()), jwtFound, err
}

func jwtValidate(req *http.Request) (jwt.JWT, bool, error) {
	jwt, err := jws.ParseJWTFromRequest(req)
	if jwt == nil {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	err = jwt.Validate(rsaPublic, signingMethod)
	if err != nil {
		switch err {
		case jws.ErrNoTokenInRequest:
			return nil, false, nil
		case jws.ErrIsNotJWT:
			return nil, true, err
		}
	}
	return jwt, true, err
}

func checkRSAKeys() error {
	if rsaPrivate == nil || rsaPublic == nil {
		return errors.New("Could not find RSA keys")
	}
	return nil
}
