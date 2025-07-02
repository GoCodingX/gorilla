//go:generate go run go.uber.org/mock/mockgen -package=repositorytest -source=repository.go -destination=repositorytest/repository.go .

package repository

import (
	"context"
)

type Repository interface {
	CreateQuote(ctx context.Context, quote *Quote) error
}
