package domain

type AccountRepository interface {
	CreateAccount(account *Account) error
	GetByApiKey(apiKey string) (*Account, error)
	GetByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}

type InvoiceRepository interface {
	CreateInvoice(invoice *Invoice) error
	GetByID(id string) (*Invoice, error)
	GetByAccountID(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
