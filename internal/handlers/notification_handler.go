package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/GoCodingX/gorilla/internal/clients"
	"github.com/GoCodingX/gorilla/internal/logger"
)

type PaymentsService struct {
	featureFlagClient clients.FeatureFlagClient
}

type NewPaymentsServiceParams struct {
	FeatureFlagClient clients.FeatureFlagClient
}

type NotificationRequestPayload struct {
	SignedPayload string `json:"signedPayload"`
}

func NewPaymentsService(params *NewPaymentsServiceParams) *PaymentsService {
	return &PaymentsService{
		featureFlagClient: params.FeatureFlagClient,
	}
}

func (s *PaymentsService) HandleNotification(w http.ResponseWriter, r *http.Request) {
	// read request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "failed to read request body"
		logger.Error(msg, slog.String("error", err.Error()))
		http.Error(w, msg, http.StatusBadRequest)

		return
	}
	defer r.Body.Close()

	// parse notification
	var notification NotificationRequestPayload
	if err := json.Unmarshal(reqBody, &notification); err != nil {
		msg := "invalid json payload"
		logger.Error(msg, slog.String("error", err.Error()))
		http.Error(w, msg, http.StatusBadRequest)

		return
	}

	// update customer feature flag
	if err := s.featureFlagClient.UnlockFeatureFlag(r.Context(), &clients.UnlockFeatureFlagParams{
		Key:    "an_awesome_feature",
		UserId: "a_user_id",
	}); err != nil {
		msg := "failed to update feature flag"
		logger.Error(msg, slog.String("error", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
