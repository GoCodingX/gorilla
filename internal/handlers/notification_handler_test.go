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
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandleNotification_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSignedPayload := "valid.jwt.token"
	mockFeatureFlagClient := clientstest.NewMockFeatureFlagClient(ctrl)

	mockFeatureFlagClient.EXPECT().UnlockFeatureFlag(gomock.Any(), &clients.UnlockFeatureFlagParams{
		Key:    "an_awesome_feature",
		UserId: "a_user_id",
	}).Return(errors.New("some terrible error happened"))

	service := handlers.NewPaymentsService(&handlers.NewPaymentsServiceParams{
		FeatureFlagClient: mockFeatureFlagClient,
	})

	requestPayload := handlers.NotificationRequestPayload{
		SignedPayload: mockSignedPayload,
	}
	reqBody, err := json.Marshal(requestPayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/v1/notifications", bytes.NewReader(reqBody))
	rec := httptest.NewRecorder()

	// act
	service.HandleNotification(rec, req)

	// assert
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestHandleNotification_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFeatureFlagClient := clientstest.NewMockFeatureFlagClient(ctrl)

	mockSignedPayload := "valid.jwt.token"

	mockFeatureFlagClient.EXPECT().UnlockFeatureFlag(gomock.Any(), &clients.UnlockFeatureFlagParams{
		Key:    "an_awesome_feature",
		UserId: "a_user_id",
	}).Return(nil)

	service := handlers.NewPaymentsService(&handlers.NewPaymentsServiceParams{
		FeatureFlagClient: mockFeatureFlagClient,
	})

	requestPayload := handlers.NotificationRequestPayload{
		SignedPayload: mockSignedPayload,
	}
	reqBody, err := json.Marshal(requestPayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/v1/notifications", bytes.NewReader(reqBody))
	rec := httptest.NewRecorder()

	// act
	service.HandleNotification(rec, req)

	// assert
	assert.Equal(t, http.StatusOK, rec.Code)
}
