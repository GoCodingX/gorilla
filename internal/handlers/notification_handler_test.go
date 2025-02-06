package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoCodingX/gorilla/internal/clients"
	"github.com/GoCodingX/gorilla/internal/clients/clientstest"
	"github.com/GoCodingX/gorilla/internal/handlers"
	"github.com/GoCodingX/gorilla/internal/jwtprocessor"
	"github.com/GoCodingX/gorilla/internal/jwtprocessor/jwtprocessortest"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandleAppleNotification_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockJwtProcessor := jwtprocessortest.NewMockProvider(ctrl)

	mockSignedPayload := "valid.jwt.token"

	mockJwtProcessor.EXPECT().ValidateAndDecodeAppleJWT(mockSignedPayload).
		Return(nil, errors.New("some terrible error happened"))

	service := handlers.NewPaymentsService(&handlers.NewPaymentsServiceParams{
		JwtProcessor: mockJwtProcessor,
	})

	requestPayload := handlers.AppleNotificationRequestPayload{
		SignedPayload: mockSignedPayload,
	}
	reqBody, err := json.Marshal(requestPayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/v1/apple-notifications", bytes.NewReader(reqBody))
	rec := httptest.NewRecorder()

	// act
	service.HandleAppleNotification(rec, req)

	// assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleAppleNotification_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFeatureFlagClient := clientstest.NewMockFeatureFlagClient(ctrl)
	mockJwtProcessor := jwtprocessortest.NewMockProvider(ctrl)

	mockSignedPayload := "valid.jwt.token"
	userId := gofakeit.UUID()

	mockFeatureFlagClient.EXPECT().UnlockFeatureFlag(gomock.Any(), &clients.UnlockFeatureFlagParams{
		Key:    "an_awesome_feature",
		UserId: userId,
	}).Return(nil)

	mockJwtProcessor.EXPECT().ValidateAndDecodeAppleJWT(mockSignedPayload).
		Return(&jwtprocessor.DecodedPayload{
			NotificationType: "SUBSCRIBED",
			ExternalId:       userId,
		}, nil)

	service := handlers.NewPaymentsService(&handlers.NewPaymentsServiceParams{
		FeatureFlagClient: mockFeatureFlagClient,
		JwtProcessor:      mockJwtProcessor,
	})

	requestPayload := handlers.AppleNotificationRequestPayload{
		SignedPayload: mockSignedPayload,
	}
	reqBody, err := json.Marshal(requestPayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/v1/apple-notifications", bytes.NewReader(reqBody))
	rec := httptest.NewRecorder()

	// act
	service.HandleAppleNotification(rec, req)

	// assert
	assert.Equal(t, http.StatusOK, rec.Code)
}
