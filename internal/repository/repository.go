package clients

import (
	"context"
)

//go:generate go run go.uber.org/mock/mockgen -package=repositorytest -source=repository.go -destination=repositorytest/repository.go .

// Repository defines the interface for feature flag service related operations.
type Repository interface {
	CreateQuote(ctx context.Context, quote *Quote) error
}
