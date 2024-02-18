package factory

import (
	"github.com/brunobrolesi/marmota-de-briga/internal/business/usecase"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/api/handler"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/repository"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/repository/config"
)

func GetBankStatementFactory() handler.Handler {
	session := config.GetScyllaSession(config.KEYSPACE)
	clientRepository := repository.NewClientRepository(session)
	transactionRepository := repository.NewTransactionRepository(session)
	uc := usecase.NewGetBankStatementUseCase(clientRepository, transactionRepository)
	h := handler.NewGetBankStatementHandler(uc)
	return h
}
