package handler

import (
	"strconv"

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
		return err
	}

	i := &usecase.InputGetBankStatement{
		ClientID: clientID,
	}

	response, err := h.getBankStatementUseCase.Execute(c.Context(), i)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
