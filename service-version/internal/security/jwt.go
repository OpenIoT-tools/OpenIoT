package security

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtToken struct {
	asymmetricKey, symmetrical string
}

func NewJWT(asymmetricKey, symmetrical string) (*jwtToken, error) {
	asynkey := os.Getenv(asymmetricKey)
	if asynkey == "" {
		return nil, fmt.Errorf("cannot find asymmetric private key")
	}
	synkey := os.Getenv(symmetrical)
	if synkey == "" {
		return nil, fmt.Errorf("cannot find symmetrical private key")
	}

	return &jwtToken{
		asymmetricKey: asynkey,
		symmetrical:   synkey,
	}, nil
}

// GenerateToken should be used to generate an rsa jwt token
func (j *jwtToken) GenerateAsymmetricToken(tokenData map[string]any, minutesLong int) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * time.Duration(minutesLong)).Unix(),
	}
	for key, value := range tokenData {
		claims[key] = value
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(j.asymmetricKey))
	if err != nil {
		return "", fmt.Errorf("cannot parse private key")
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
}

func (j *jwtToken) ValidateSymmetricalToken(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.symmetrical), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	if expirationTime.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
