package factory

import "github.com/marcos-dev/imersao5-gateway/domain/repository"

// Cria uma Fabrica de repositório da transação
type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
