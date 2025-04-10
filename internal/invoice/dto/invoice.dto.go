package dto

import (
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

// CreateInvoiceDTO is a struct that represents a create invoice request
type CreateInvoiceDTO struct {
	ApiKey          string
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	PaymentType     string  `json:"payment_type"`
	CardNumber      string  `json:"card_number"`
	CardholderName  string  `json:"cardholder_name"`
	CVV             string  `json:"cvv"`
	ExpirationMonth string  `json:"expiration_month"`
	ExpirationYear  string  `json:"expiration_year"`
}

// InvoiceResponseDTO is a struct that represents an invoice response
type InvoiceResponseDTO struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ToInvoiceDTO converts a Invoice to an InvoiceResponseDTO
func ToInvoiceDTO(input *CreateInvoiceDTO, accountID string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		CardNumber:      input.CardNumber,
		CardholderName:  input.CardholderName,
		CVV:             input.CVV,
		ExpirationMonth: input.ExpirationMonth,
		ExpirationYear:  input.ExpirationYear,
	}

	return domain.NewInvoice(
		accountID,
		input.Amount,
		input.Description,
		input.PaymentType,
		card,
	)
}

// FromInvoiceDTO converts an Invoice to an InvoiceResponseDTO
func FromInvoiceDTO(invoice *domain.Invoice) *InvoiceResponseDTO {
	return &InvoiceResponseDTO{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
