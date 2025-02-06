package apple

import (
	"crypto/x509"
	"encoding/json"
	"fmt"

	"github.com/GoCodingX/gorilla/internal/jwtprocessor"
	"github.com/GoCodingX/gorilla/internal/utils"
)

type JwtProcessor struct {
	rootCert *x509.Certificate
}

func NewJwtProcessor(rootCert *x509.Certificate) *JwtProcessor {
	return &JwtProcessor{
		rootCert: rootCert,
	}
}

func (p *JwtProcessor) ValidateAndDecodeAppleJWT(signedPayload string) (*jwtprocessor.DecodedPayload, error) {
	payload, err := utils.ValidateJwtPayloadAndExtractPayload(signedPayload, p.rootCert)
	if err != nil {
		return nil, fmt.Errorf("failed to validate jwt payload: %w", err)
	}

	var decodedPayload jwtprocessor.DecodedPayload

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	if err := json.Unmarshal(payloadJSON, &decodedPayload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	return &decodedPayload, nil
}
