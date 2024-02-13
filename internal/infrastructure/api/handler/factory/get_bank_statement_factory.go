package factory

import (
	"github.com/brunobrolesi/marmota-de-briga/internal/business/usecase"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/api/handler"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/repository"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/repository/config"
)

func GetBankStatementFactory() handler.Handler {
	config := config.NewScyllaClient(config.KEYSPACE)
	clientRepository := repository.NewClientRepository(config)
	transactionRepository := repository.NewTransactionRepository(config)
	uc := usecase.NewGetBankStatementUseCase(clientRepository, transactionRepository)
	h := handler.NewGetBankStatementHandler(uc)
	return h
}
