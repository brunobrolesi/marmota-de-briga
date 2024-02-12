package factory

import (
	"github.com/brunobrolesi/marmota-de-briga/internal/business/usecase"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/api/handler"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/repository"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/repository/config"
)

func CreateTransactionFactory() handler.Handler {
	config := config.NewScyllaClient(config.KEYSPACE)
	clientRepository := repository.NewClientRepository(config)
	transactionRepository := repository.NewTransactionRepository(config)
	uc := usecase.NewCreateTransactionUseCase(clientRepository, transactionRepository)
	h := handler.NewCreateTransactionHandler(uc)
	return h
}
