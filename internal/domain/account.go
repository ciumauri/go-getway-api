package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Account is a struct that represents an account
type Account struct {
	ID        string
	Name      string
	Email     string
	ApiKey    string
	Balance   float64
	mu        sync.Mutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

// generateApiKey generates a random API key
func generateApiKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// NewAccount creates a new account
func NewAccount(name, email string) *Account {

	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		ApiKey:    generateApiKey(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

// AddBalance adds a balance to the account
func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	a.UpdatedAt = time.Now()
}
