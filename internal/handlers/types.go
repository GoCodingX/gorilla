package handlers

import (
	"time"

	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/google/uuid"
)

func toRepoAuthor(author *openapi.CreateAuthorRequest) *repository.Author {
	return &repository.Author{
		ID:        uuid.New(),
		Name:      author.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func toRepoQuote(quote *openapi.CreateQuoteRequest, username string) *repository.Quote {
	return &repository.Quote{
		ID:              uuid.New(),
		Text:            quote.Text,
		AuthorID:        quote.AuthorId,
		CreatorUsername: username,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}
}
