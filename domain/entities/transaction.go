package entities

type Transaction struct {
	Value       int
	Type        TransactionType // "c" (credit) or "d" (debit)
	Description string
}

type TransactionType = string

const (
	Credit TransactionType = "c"
	Debit  TransactionType = "d"
)
