package handlers

import (
	"github.com/GoCodingX/gorilla/internal/repository"
)

type QuotesService struct {
	repo repository.Repository
}

type NewQuotesServiceParams struct {
	Repo repository.Repository
}

func NewQuotesService(params *NewQuotesServiceParams) *QuotesService {
	return &QuotesService{
		repo: params.Repo,
	}
}
