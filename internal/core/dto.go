package core

type CreateTransactionInput struct {
	Amount      int    `json:"valor"`
	Operation   string `json:"tipo"`
	Description string `json:"descricao"`
}

type CreateTransactionOutput struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type Balance struct {
	Total int    `json:"total"`
	Date  string `json:"data_extrato"`
	Limit int    `json:"limite"`
}

type TransactionOutput struct {
	Amount      int    `json:"valor"`
	Operation   string `json:"tipo"`
	Description string `json:"descricao"`
	Date        string `json:"realizada_em"`
}

type GetBalanceOutput struct {
	Balance      Balance             `json:"saldo"`
	Transactions []TransactionOutput `json:"ultimas_transacoes"`
}
