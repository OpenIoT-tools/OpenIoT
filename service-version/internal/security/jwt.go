package security

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
}

func (j *Jwt) loadPrivateKey(privateKeyName string) (*rsa.PrivateKey, error) {
	pem := os.Getenv(privateKeyName)
	if pem == "" {
		return nil, fmt.Errorf("cannot find private key")
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pem))
	if err != nil {
		return nil, fmt.Errorf("cannot parse private key")
	}
	return key, nil
}

// GenerateToken should be used to generate an asynchronous jwt token
func (j *Jwt) GenerateToken(tokenData map[string]any, minutesLong int, privateKeyName string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * time.Duration(minutesLong)).Unix(),
	}
	for key, value := range tokenData {
		claims[key] = value
	}

	privateKey, err := j.loadPrivateKey(privateKeyName)
	if err != nil {
		return "", err
	}
	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
}
