package handler_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/api/handler"
	mock_usecase "github.com/brunobrolesi/marmota-de-briga/mocks/internal_/business/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionHandler(t *testing.T) {
	type TestSuite struct {
		App                      *fiber.App
		CreateTransactionUseCase *mock_usecase.MockCreateTransactionUseCase
	}

	setup := func(t *testing.T) TestSuite {
		app := fiber.New()
		uc := mock_usecase.NewMockCreateTransactionUseCase(t)
		h := handler.NewCreateTransactionHandler(uc)
		app.Post("/clientes/:id/transacoes", h.Handle)
		return TestSuite{
			App:                      app,
			CreateTransactionUseCase: uc,
		}
	}

	makeRequestBody := func() io.Reader {
		b := []byte(`{"valor": 100, "tipo": "c", "descricao": "test"}`)
		return bytes.NewReader(b)
	}
	t.Run("should return an error if the request body is invalid", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("POST", "/clientes/1/transacoes", nil)

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, fiber.ErrUnprocessableEntity.Message, string(body))
	})

	t.Run("should return an 404 if the client id is invalid", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("POST", "/clientes/a/transacoes", makeRequestBody())
		req.Header.Set("Content-Type", "application/json")

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
		req := httptest.NewRequest("POST", "/clientes/0/transacoes", makeRequestBody())
		req.Header.Set("Content-Type", "application/json")

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"message":"client not found"}`, string(body))
	})

	t.Run("should return an 404 if the client id is not registered in db", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("POST", "/clientes/1/transacoes", makeRequestBody())
		req.Header.Set("Content-Type", "application/json")

		testSuite.CreateTransactionUseCase.On("Execute", mock.Anything, mock.Anything).Return(nil, model.ErrClientNotFound).Once()

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"message":"client not found"}`, string(body))
	})

	t.Run("should return an 422 if transaction exceed client limit", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("POST", "/clientes/1/transacoes", makeRequestBody())
		req.Header.Set("Content-Type", "application/json")

		testSuite.CreateTransactionUseCase.On("Execute", mock.Anything, mock.Anything).Return(nil, model.ErrClientLimitExceeded).Once()

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"message":"client limit exceeded"}`, string(body))
	})

	t.Run("should return an 500 if use case return an unmapped error", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("POST", "/clientes/1/transacoes", makeRequestBody())
		req.Header.Set("Content-Type", "application/json")

		testSuite.CreateTransactionUseCase.On("Execute", mock.Anything, mock.Anything).Return(nil, model.ErrInternalServerError).Once()

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"message":"internal server error"}`, string(body))
	})

	t.Run("should return an 200, client balance and limit on success", func(t *testing.T) {
		testSuite := setup(t)

		// http.Request
		req := httptest.NewRequest("POST", "/clientes/1/transacoes", makeRequestBody())
		req.Header.Set("Content-Type", "application/json")

		testSuite.CreateTransactionUseCase.On("Execute", mock.Anything, mock.Anything).Return(&model.Client{
			ID:             1,
			AccountBalance: 100,
			AccountLimit:   1000,
		}, nil).Once()

		// http.Response
		resp, _ := testSuite.App.Test(req)

		// Asserts
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, `{"limite":1000,"saldo":100}`, string(body))
	})
}
