package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/GoCodingX/gorilla/internal/clients"
	"github.com/GoCodingX/gorilla/internal/jwtprocessor"
	"github.com/GoCodingX/gorilla/internal/logger"
)

type PaymentsService struct {
	featureFlagClient clients.FeatureFlagClient
	jwtProcessor      jwtprocessor.Provider
}

type NewPaymentsServiceParams struct {
	FeatureFlagClient clients.FeatureFlagClient
	JwtProcessor      jwtprocessor.Provider
}

type AppleNotificationRequestPayload struct {
	SignedPayload string `json:"signedPayload"`
}

func NewPaymentsService(params *NewPaymentsServiceParams) *PaymentsService {
	return &PaymentsService{
		featureFlagClient: params.FeatureFlagClient,
		jwtProcessor:      params.JwtProcessor,
	}
}

func (s *PaymentsService) HandleAppleNotification(w http.ResponseWriter, r *http.Request) {
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
	var notification AppleNotificationRequestPayload
	if err := json.Unmarshal(reqBody, &notification); err != nil {
		msg := "invalid json payload"
		logger.Error(msg, slog.String("error", err.Error()))
		http.Error(w, msg, http.StatusBadRequest)

		return
	}

	// verify signature and decode payload
	decodedPayload, err := s.jwtProcessor.ValidateAndDecodeAppleJWT(notification.SignedPayload)
	if err != nil {
		msg := "jwt validation failed"
		logger.Error(msg, slog.String("error", err.Error()))
		http.Error(w, msg, http.StatusBadRequest)

		return
	}

	// enable feature flag if the customer has subscribed (simplified logic for the sake of exercise)
	if decodedPayload.NotificationType == "SUBSCRIBED" {
		// update customer feature flag
		if err := s.featureFlagClient.UnlockFeatureFlag(r.Context(), &clients.UnlockFeatureFlagParams{
			Key:    "an_awesome_feature",
			UserId: decodedPayload.ExternalId,
		}); err != nil {
			msg := "failed to update feature flag"
			logger.Error(msg, slog.String("error", err.Error()))
			http.Error(w, msg, http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
