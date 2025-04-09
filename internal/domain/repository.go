package domain

type AccountRepository interface {
	CreateAccount(account *Account) error
	GetByApiKey(apiKey string) (*Account, error)
	GetByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}
