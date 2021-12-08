package factory

import "github.com/marcosaraujo-dev/Imersao_FullCycle/domain/repository"

// Cria uma Fabrica de repositório da transação
type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
