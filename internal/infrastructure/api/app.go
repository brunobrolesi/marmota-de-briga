package api

import (
	"log"

	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/api/handler/factory"
	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New()

	h := factory.CreateTransactionFactory()

	app.Post("/clientes/:id/transacoes", h.Handle)
	app.Get("/clientes/:id/extrato", nil)

	log.Fatal(app.Listen(":8080"))
}
