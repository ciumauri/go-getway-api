package dto

import (
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

// CreateAccountDTO is a struct that represents a create account request
type CreateAccountDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// AccountResponseDTO is a struct that represents an account response
type AccountResponseDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ApiKey    string    `json:"api_key"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToAccountDTO converts a CreateAccountDTO to an Account
func ToAccountDTO(input CreateAccountDTO) *domain.Account {
	return domain.NewAccount(
		input.Name,
		input.Email,
	)
}

// FromAccountDTO converts an Account to an AccountResponseDTO
func FromAccountDTO(account *domain.Account) AccountResponseDTO {
	return AccountResponseDTO{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		ApiKey:    account.ApiKey,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
