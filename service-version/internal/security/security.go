package security

type SecurityToken interface {
	GenerateToken(tokenData map[string]any, minutesLong int, privateKeyName string) (string, error)
}
