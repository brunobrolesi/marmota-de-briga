package handler

import (
	"errors"
	"strconv"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateTransactionRequestBody struct {
	Value       int    `json:"valor" validate:"required,min=1"`
	Type        string `json:"tipo" validate:"required,oneof=c d"`
	Description string `json:"descricao" validate:"required,min=1,max=10"`
}

type CreateTransactionResponseBody struct {
	Limit   model.MonetaryValue `json:"limite"`
	Balance model.ClientBalance `json:"saldo"`
}

type createTransactionHandler struct {
	createTransactionUsecase usecase.CreateTransactionUseCase
	validator                *validator.Validate
}

func NewCreateTransactionHandler(
	createTransactionUsecase usecase.CreateTransactionUseCase,
	validator *validator.Validate,
) Handler {
	return &createTransactionHandler{
		createTransactionUsecase: createTransactionUsecase,
		validator:                validator,
	}
}

func (h *createTransactionHandler) Handle(c *fiber.Ctx) error {
	body := new(CreateTransactionRequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": fiber.ErrUnprocessableEntity.Message})
	}

	if err := h.validator.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": fiber.ErrUnprocessableEntity.Message})
	}

	clientID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": model.ErrClientNotFound.Error()})
	}

	if clientID <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": model.ErrClientNotFound.Error()})
	}

	i := &usecase.InputCreateTransaction{
		ClientID:    clientID,
		Value:       model.MonetaryValue(body.Value),
		Type:        model.TransactionType(body.Type),
		Description: body.Description,
	}

	client, err := h.createTransactionUsecase.Execute(c.Context(), i)

	if errors.Is(err, model.ErrClientNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	if errors.Is(err, model.ErrClientLimitExceeded) {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := &CreateTransactionResponseBody{
		Limit:   client.AccountLimit,
		Balance: client.AccountBalance,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
