package handlers_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GoCodingX/gorilla/internal/handlers"
	"github.com/GoCodingX/gorilla/internal/repository/repositorytest"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

type newEchoContextParams struct {
	method  string
	target  string
	payload string
}

// newEchoContext returns an echo.Context
func newEchoContext(params *newEchoContextParams) (echo.Context, *httptest.ResponseRecorder) {
	var (
		method string
		target string
		body   *strings.Reader
	)

	if params != nil {
		if params.method != "" {
			method = params.method
		}

		if params.target != "" {
			target = params.target
		}

		if params.payload != "" {
			body = strings.NewReader(params.payload)
		}
	}

	e := echo.New()
	req := httptest.NewRequest(method, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}

func newServiceWithMockRepo(t *testing.T) (*handlers.QuotesService, *repositorytest.MockRepository) {
	ctrl := gomock.NewController(t)
	repo := repositorytest.NewMockRepository(ctrl)
	svc := handlers.NewQuotesService(&handlers.NewQuotesServiceParams{
		Repo: repo,
	})

	return svc, repo
}
