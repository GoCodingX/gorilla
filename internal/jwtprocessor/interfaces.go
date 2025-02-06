package jwtprocessor

//go:generate go run go.uber.org/mock/mockgen -package=jwtprocessortest -source=interfaces.go -destination=jwtprocessortest/jwt.go .

type DecodedPayload struct {
	NotificationType string `json:"notificationType"`
	ExternalId       string `json:"external_id"`
}

type Provider interface {
	ValidateAndDecodeAppleJWT(signedPayload string) (*DecodedPayload, error)
}
