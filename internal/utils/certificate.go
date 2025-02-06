package utils

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func ParsePemCert(pemCert string) (*x509.Certificate, error) {
	// clean the certificate content
	pemCert = strings.TrimSpace(pemCert)

	// decode the PEM block
	block, _ := pem.Decode([]byte(pemCert))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	// parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %w", err)
	}

	return cert, nil
}

func ValidateJwtPayloadAndExtractPayload(signedPayload string, cert *x509.Certificate) (jwt.Claims, error) {
	// parse the JWT(S) token
	token, err := jwt.Parse(signedPayload, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return cert.PublicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("jwt validation failed: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid JWT")
	}

	return token.Claims, nil
}
