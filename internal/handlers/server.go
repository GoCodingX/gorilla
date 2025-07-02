package handlers

import (
	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/GoCodingX/gorilla/internal/repository/pg"
)

type QuotesService struct {
	repo repository.Repository
}

type NewQuotesServiceParams struct {
	Repo *pg.Repository
}

func NewQuotesService(params *NewQuotesServiceParams) *QuotesService {
	return &QuotesService{
		repo: params.Repo,
	}
}
