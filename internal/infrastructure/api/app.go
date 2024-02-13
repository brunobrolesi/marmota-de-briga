package api

import (
	"log"

	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/api/handler/factory"
	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New()

	handlerCreateTransaction := factory.CreateTransactionFactory()
	handlerGetBankStatement := factory.GetBankStatementFactory()

	app.Post("/clientes/:id/transacoes", handlerCreateTransaction.Handle)
	app.Get("/clientes/:id/extrato", handlerGetBankStatement.Handle)

	log.Fatal(app.Listen(":8080"))
}
