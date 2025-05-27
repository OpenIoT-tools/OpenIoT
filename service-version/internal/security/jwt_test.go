package security

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewJWT(t *testing.T) {
	t.Run("when asymmetric env var is missing, should return error", func(t *testing.T) {
		os.Unsetenv("ASYM")
		os.Setenv("SYM", "foo")
		_, err := NewJWT("ASYM", "SYM")
		assert.EqualError(t, err, "cannot find asymmetric private key")
	})

	t.Run("when symmetrical env var is missing, should return error", func(t *testing.T) {
		os.Setenv("ASYM", "bar")
		os.Unsetenv("SYM")
		_, err := NewJWT("ASYM", "SYM")
		assert.EqualError(t, err, "cannot find symmetrical private key")
	})

	t.Run("when both env vars are present, should succeed", func(t *testing.T) {
		os.Setenv("ASYM", "bar")
		os.Setenv("SYM", "foo")
		j, err := NewJWT("ASYM", "SYM")
		assert.NoError(t, err)
		assert.Equal(t, "bar", j.asymmetricKey)
		assert.Equal(t, "foo", j.symmetrical)
	})
}

func TestGenerateAsymmetricToken(t *testing.T) {
	os.Setenv("ASYM", os.Getenv("RSA_PRIV_TEST"))
	os.Setenv("SYM", os.Getenv("RSA_PRIV_TEST"))
	j, _ := NewJWT("ASYM", "SYM")

	t.Run("when generating a token with valid claims, should return signed token", func(t *testing.T) {
		claims := map[string]any{"foo": "bar"}
		tokenStr, err := j.GenerateAsymmetricToken(claims, 5)
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenStr)

		parsed, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("testRSAPub")))
		})
		assert.NoError(t, err)
		assert.True(t, parsed.Valid)
	})
}

func TestValidateSymmetricalToken(t *testing.T) {
	os.Setenv("SYM", "mysecret")
	os.Setenv("ASYM", "mysecret")
	j, _ := NewJWT("ASYM", "SYM")

	rawValid := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "baz",
		"exp": time.Now().Add(1 * time.Minute).Unix(),
	})
	validStr, _ := rawValid.SignedString([]byte("mysecret"))

	t.Run("when token is valid, should return claims", func(t *testing.T) {
		out, err := j.ValidateSymmetricalToken(validStr)
		assert.NoError(t, err)
		assert.Equal(t, "baz", out["foo"])
	})

	t.Run("when token signature is invalid, should return error", func(t *testing.T) {
		_, err := j.ValidateSymmetricalToken(validStr + "x")
		assert.Error(t, err)
	})

	t.Run("when token is expired, should return 'invalid token' error", func(t *testing.T) {
		rawExpired := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"foo": "baz",
			"exp": time.Now().Add(-1 * time.Minute).Unix(),
		})
		expStr, _ := rawExpired.SignedString([]byte("mysecret"))

		_, err := j.ValidateSymmetricalToken(expStr)
		assert.EqualError(t, err, "invalid token")
	})
}
