package helper

import (
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTPayload struct {
	User map[string]interface{} `json:"user"`
	jwt.RegisteredClaims
}

func CreateToken(dir string, data map[string]interface{}) (string, error) {
	privKey, err := ioutil.ReadFile(dir)
	if err != nil {
		return "", nil
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		return "", nil
	}

	claims := JWTPayload{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", nil
	}

	return token, nil
}
