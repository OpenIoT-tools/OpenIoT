package security

type SecurityToken interface {
	GenerateAsymmetricToken(tokenData map[string]any, minutesLong int) (string, error)
	ValidateSymmetricalToken(tokenStr string) (map[string]interface{}, error)
}
