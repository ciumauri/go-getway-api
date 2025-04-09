package repository

import (
	"database/sql"
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

// AccountRepository is a repository for accounts
type AccountRepository struct {
	db *sql.DB
}

// NewAccountRepository creates a new account repository
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// CreateAccount creates a new account in the database
func (r *AccountRepository) CreateAccount(account *domain.Account) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) 
	  	VALUES ($1, $2, $3, $4, $5, $6, $7)	  
	  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.ApiKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetByApiKey retrieves an account by its API key
func (r *AccountRepository) GetByApiKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at 
		FROM accounts 
		WHERE api_key = $1
	`, apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

// GetByID retrieves an account by its ID
func (r *AccountRepository) GetByID(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at 
		FROM accounts 
		WHERE id = $1
	 `, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

// UpdateBalance updates the balance of an account
func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow(`
		SELECT balance FROM accounts 
		WHERE id = $1
		FOR UPDATE
	`, account.ID).Scan(&currentBalance)

	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE accounts
		SET balance = $1, updated_at = $2
		WHERE id = $3
	`, account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
