package clients

import (
	"context"
)

//go:generate go run go.uber.org/mock/mockgen -package=clientstest -source=interfaces.go -destination=clientstest/clients.go .

type UnlockFeatureFlagParams struct {
	Key    string
	UserId string
}

// FeatureFlagClient defines the interface for feature flag service related operations.
type FeatureFlagClient interface {
	UnlockFeatureFlag(ctx context.Context, params *UnlockFeatureFlagParams) error
}
