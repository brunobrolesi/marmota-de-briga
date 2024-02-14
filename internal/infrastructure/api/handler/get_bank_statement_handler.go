package handler

import (
	"errors"
	"strconv"

	"github.com/brunobrolesi/marmota-de-briga/internal/business/model"
	"github.com/brunobrolesi/marmota-de-briga/internal/business/usecase"
	"github.com/gofiber/fiber/v2"
)

type getBankStatementHandler struct {
	getBankStatementUseCase usecase.GetBankStatementUseCase
}

func NewGetBankStatementHandler(getBankStatementUseCase usecase.GetBankStatementUseCase) Handler {
	return &getBankStatementHandler{
		getBankStatementUseCase: getBankStatementUseCase,
	}
}

func (h *getBankStatementHandler) Handle(c *fiber.Ctx) error {
	clientID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": model.ErrClientNotFound.Error()})
	}

	if clientID <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": model.ErrClientNotFound.Error()})
	}

	i := &usecase.InputGetBankStatement{
		ClientID: clientID,
	}

	response, err := h.getBankStatementUseCase.Execute(c.Context(), i)

	if errors.Is(err, model.ErrClientNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
