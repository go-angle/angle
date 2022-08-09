package secure

import (
	"errors"

	"github.com/go-angle/angle/config"
	"github.com/go-angle/angle/di"
	"github.com/golang-jwt/jwt/v4"
)

// jwtSign Json Web Token
type jwtSign struct {
	key []byte
}

func newJWTSign(config *config.Config) Signer {
	return &jwtSign{
		key: []byte(config.SecretKey),
	}
}

func (j *jwtSign) Sign(claims SignClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString(j.key)
}

func (j *jwtSign) Validate(tokenStr string) (SignClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return SignClaims(claims), nil
	}
	return nil, errors.New("invalid token")
}

func init() {
	di.Provide(newJWTSign)
}
