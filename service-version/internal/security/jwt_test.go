package security

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const testRSAPriv = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDYm1hfeQh/CBrd
uDGtP8vzYXGj39AyP0rEe9dal2q8osQ82L1dYINBfkJLgSQWY6WtLZ7eKHXyqVq0
L0QO7iEusbv0UtvCaR7JzRqku1gpfa8y5IJpzCsE8e+8os2EP/R878Zy/q3USQri
ESXl8ZA9/mwdb7dYgGfq0gQNs1l6f/FY2Ge2zkN+YVtHXSImIqsquv3U3ppf2b/G
twJoa7AGxf8oyQq42JqHPPiWBOIwbISs7b9//BZajd47e0ig9H+XIno3FiF788mU
0uC6GZgYtuK3rHcedshgStWaOWGq33PzSD7F1Cku0rs+Zy4v2OTKlakvERjqnAI6
DrhvsICDAgMBAAECggEANSXUxsRLXeQOt4V5BvIG6IZp3FRP4Oxfu5tJaYyKn5Nu
hCGIQU+rlRfWS3F5+dZkcfwmJYuW+T4tp9WLJQCv+qWYoSftjSO+7rQZS/898BxL
LzkbeLskeRemA8qMk90fb8Jjlpa/7z0m4vZjLqsZgeSaAZUGTzxegyMZ5ISzuVcU
t3GIFM5lC7Zhq9O0eCvYsoF9zNpF+jBSolgxfCXbeNhjNILSV5UP/Vxbd5Hig1Wr
SBOEsXfhBH9kYhWgrentsQPlKdO1S3s/4AZtIDnXaLzYRJjjGvpty/XaPG5nvO90
zND4O6IJ8ebrutms9kcr9oxERp7EEttaRT7NMPXW1QKBgQDxoRZPP93Lmu+RBZz1
ZNjnFoYx8ntXLNLicsTCT1pg96hzXb490CK2GhKtJsaDm4JPeFHa5MDaLhhvfZE+
aRjR0UTSNMrP4EKmQuH8wPLulArkDTvgOGB4/CvgEwu5sQSPJaLc93LSG0wcU0tg
w5Ay6eJiwZ+f2L78ZFbj53ZsTwKBgQDlfUfKX0JEYsfvPEa1joXDHj2giXM1gDe3
tN6RcXUZz8vWG936828gjp+SRXTMXKKXHfE0FjIQvwLFeU6mRbsIiofuxwCpE7ow
0uEip8r6MOvRTY17GMiVs/8l42BPgvphLFddpDhSPtqgxMyCJX514U01G8P5iJJA
wLn/DthXjQKBgQCyISlLPzAJe60E4nh1x+THZ6mChNo8yASngChhfxKHzcvXUNKA
y0HrsYj3MVDDkQ8d+vuNPXqW3ciR2KYMtnnfpEDyHo0tdoxr+X6UUMJG88ibms37
XFRynKTTyfao7EEKLprD/4AcOJeBgj5dlY5mmlmUqF0/ABC8DFSqKub4rQKBgDvZ
xJNe7KVxLbuS+M348aHFjUUFE+rsGZvrq/A6qtOkggDBG8+LFXxkNfGr854ouq62
vGVdNIlxh4OA86YhoXnEx1bax61Q28pH9TdYk3NUtuvLUg68k/OXEgALoN6bxjj5
m39siNPKWLJ9IOkAZk7QB5bGOwIlQB9rJBcFBB8xAoGAFyD80BhbZISlkzE+BgX2
7MsrAzeFIJS/1BhH9GZpf47lmrHc6UkjHjAJDDysUXrO31GvwywafK5utyiorR1+
TfoK8QTJyq83L/kn19iRjk+SWYirQY5AIp1sjp/VnfcBuOJ2H16vWbGdVibRbcFj
29KQn59ZgBP5H50ZJrBXPNQ=
-----END PRIVATE KEY-----`

const testRSAPub = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2JtYX3kIfwga3bgxrT/L
82Fxo9/QMj9KxHvXWpdqvKLEPNi9XWCDQX5CS4EkFmOlrS2e3ih18qlatC9EDu4h
LrG79FLbwmkeyc0apLtYKX2vMuSCacwrBPHvvKLNhD/0fO/Gcv6t1EkK4hEl5fGQ
Pf5sHW+3WIBn6tIEDbNZen/xWNhnts5DfmFbR10iJiKrKrr91N6aX9m/xrcCaGuw
BsX/KMkKuNiahzz4lgTiMGyErO2/f/wWWo3eO3tIoPR/lyJ6NxYhe/PJlNLguhmY
GLbit6x3HnbIYErVmjlhqt9z80g+xdQpLtK7PmcuL9jkypWpLxEY6pwCOg64b7CA
gwIDAQAB
-----END PUBLIC KEY-----`

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
	os.Setenv("ASYM", testRSAPriv)
	os.Setenv("SYM", testRSAPriv)
	j, _ := NewJWT("ASYM", "SYM")

	t.Run("when generating a token with valid claims, should return signed token", func(t *testing.T) {
		claims := map[string]any{"foo": "bar"}
		tokenStr, err := j.GenerateAsymmetricToken(claims, 5)
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenStr)

		parsed, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(testRSAPub))
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
