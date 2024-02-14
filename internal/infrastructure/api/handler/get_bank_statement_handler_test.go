package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/api/handler"
	mock_usecase "github.com/brunobrolesi/marmota-de-briga/mocks/internal_/business/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBankStatementHandler(t *testing.T) {
	type TestSuite struct {
		App                     *fiber.App
		GetBankStatementUseCase *mock_usecase.MockGetBankStatementUseCase
	}

	setup := func(t *testing.T) TestSuite {
		app := fiber.New()
		uc := mock_usecase.NewMockGetBankStatementUseCase(t)
		h := handler.NewGetBankStatementHandler(uc)
		app.Get("/clientes/:id/extrato", h.Handle)
		return TestSuite{
			App:                     app,
			GetBankStatementUseCase: uc,
		}
	}

	t.Run("should return an 404 if the client id is invalid", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("GET", "/clientes/a/extrato", nil)

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"message":"client not found"}`, string(body))
	})

	t.Run("should return an 404 if the client id is equal or less than 0", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("GET", "/clientes/0/extrato", nil)

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"message":"client not found"}`, string(body))
	})

	t.Run("should return an 404 if the client id not exists in db", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("GET", "/clientes/1/extrato", nil)

		testSuite.GetBankStatementUseCase.On("Execute", mock.Anything, mock.Anything).Return(nil, model.ErrClientNotFound).Once()

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"message":"client not found"}`, string(body))
	})

	t.Run("should return an 500 if use case return an unmapped error", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("GET", "/clientes/1/extrato", nil)

		testSuite.GetBankStatementUseCase.On("Execute", mock.Anything, mock.Anything).Return(nil, model.ErrInternalServerError).Once()

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"message":"internal server error"}`, string(body))
	})

	t.Run("should return an 200 and bank statement on success", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("GET", "/clientes/1/extrato", nil)
		req.Header.Set("Content-Type", "application/json")

		timeUTC, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
		if err != nil {
			t.Fatalf("error parsing time: %v", err)
		}

		statement := &model.BankStatement{
			Balance: model.BankStatementBalance{
				Total:     100,
				CreatedAt: timeUTC,
				Limit:     1000,
			},
			Transactions: []model.BankStatementTransaction{
				{
					Value:       100,
					Type:        model.Credit,
					Description: "description",
					CreatedAt:   timeUTC,
				},
				{
					Value:       100,
					Type:        model.Debit,
					Description: "description",
					CreatedAt:   timeUTC,
				},
			},
		}

		testSuite.GetBankStatementUseCase.On("Execute", mock.Anything, mock.Anything).Return(statement, nil).Once()

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"saldo":{"total":100,"data_extrato":"2021-01-01T00:00:00Z","limite":1000},"ultimas_transacoes":[{"valor":100,"tipo":"c","descricao":"description","realizada_em":"2021-01-01T00:00:00Z"},{"valor":100,"tipo":"d","descricao":"description","realizada_em":"2021-01-01T00:00:00Z"}]}`, string(body))
	})
}
