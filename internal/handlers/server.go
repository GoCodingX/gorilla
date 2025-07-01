package handlers

import "github.com/GoCodingX/gorilla/internal/clients"

type QuotesService struct {
	featureFlagClient clients.FeatureFlagClient
}

type NewQuotesServiceParams struct {
	FeatureFlagClient clients.FeatureFlagClient
}

func NewQuotesService(params *NewQuotesServiceParams) *QuotesService {
	return &QuotesService{
		featureFlagClient: params.FeatureFlagClient,
	}
}
