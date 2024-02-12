package handler

import (
	"strconv"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/usecase"
	"github.com/gofiber/fiber/v2"
)

type CreateTransactionRequestBody struct {
	Value       model.MonetaryValue   `json:"valor"`
	Type        model.TransactionType `json:"tipo"`
	Description string                `json:"descricao"`
}

type CreateTransactionResponseBody struct {
	Limit   model.MonetaryValue `json:"limite"`
	Balance model.ClientBalance `json:"saldo"`
}

type createTransactionHandler struct {
	createTransactionUsecase usecase.CreateTransactionUseCase
}

func NewCreateTransactionHandler(createTransactionUsecase usecase.CreateTransactionUseCase) Handler {
	return &createTransactionHandler{
		createTransactionUsecase: createTransactionUsecase,
	}
}

func (h *createTransactionHandler) Handle(c *fiber.Ctx) error {
	body := new(CreateTransactionRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	clientID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	i := &usecase.InputCreateTransaction{
		ClientID:    clientID,
		Value:       body.Value,
		Type:        body.Type,
		Description: body.Description,
	}

	client, err := h.createTransactionUsecase.Execute(c.Context(), i)
	if err != nil {
		return err
	}

	response := &CreateTransactionResponseBody{
		Limit:   client.AccountLimit,
		Balance: client.AccountBalance,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
