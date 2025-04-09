package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Status is a type that represents the status of an invoice
type Status string

// StatusPending is the status of an invoice that is pending
const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

// Invoice is a struct that represents an invoice
type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         string
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// CreditCard is a struct that represents a credit card
type CreditCard struct {
	CardNumber      string
	CardholderName  string
	CVV             string
	ExpirationMonth string
	ExpirationYear  string
}

// NewInvoice creates a new invoice
func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.CardNumber[len(card.CardNumber)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         string(StatusPending),
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

// Process processes an invoice
func (i *Invoice) Process() error {
	if i.Amount <= 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	var newStatus Status

	if randomSource.Float64() < 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = string(newStatus)

	return nil
}

// UpdateStatus updates the status of an invoice
func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != string(StatusPending) {
		return ErrInvalidStatus
	}

	i.Status = string(newStatus)
	i.UpdatedAt = time.Now()

	return nil
}
