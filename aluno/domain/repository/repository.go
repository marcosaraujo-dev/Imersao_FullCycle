package repository

//  interface da Transação
type TransactionRepository interface {
	Insert(id string, account string, amount float64, status string, errorMessage string) error
}
