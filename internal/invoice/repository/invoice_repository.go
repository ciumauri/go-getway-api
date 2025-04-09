package repository

import (
	"database/sql"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

// InvoiceRepository is a repository for invoices
type InvoiceRepository struct {
	db *sql.DB
}

// NewInvoiceRepository creates a new invoice repository
func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// CreateInvoice creates a new invoice in the database
func (r *InvoiceRepository) CreateInvoice(invoice *domain.Invoice) error {
	_, err := r.db.Exec(`
		INSERT INTO invoices (id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, invoice.ID, invoice.AccountID, invoice.Amount, invoice.Status, invoice.Description, invoice.PaymentType, invoice.CardLastDigits, invoice.CreatedAt, invoice.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetByID retrieves an invoice by its ID
func (r *InvoiceRepository) GetByID(id string) (*domain.Invoice, error) {
	var invoice domain.Invoice

	err := r.db.QueryRow(`
		SELECT id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at 
		FROM invoices 
		WHERE id = $1
	`, id).Scan(
		&invoice.ID,
		&invoice.AccountID,
		&invoice.Amount,
		&invoice.Status,
		&invoice.Description,
		&invoice.PaymentType,
		&invoice.CardLastDigits,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrInvoiceNotFound
	}

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

// GetByAccountID retrieves all invoices by account ID
func (r *InvoiceRepository) GetByAccountID(accountID string) ([]*domain.Invoice, error) {
	rows, err := r.db.Query(`
		SELECT id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at 
		FROM invoices 
		WHERE account_id = $1
	`, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*domain.Invoice
	for rows.Next() {
		var invoice domain.Invoice
		err := rows.Scan(
			&invoice.ID,
			&invoice.AccountID,
			&invoice.Amount,
			&invoice.Status,
			&invoice.Description,
			&invoice.PaymentType,
			&invoice.CardLastDigits,
			&invoice.CreatedAt,
			&invoice.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

// UpdateStatus updates the status of an invoice
func (r *InvoiceRepository) UpdateStatus(invoice *domain.Invoice) error {
	rows, err := r.db.Exec(`
		UPDATE invoices 
		SET status = $1, updated_at = $2 
		WHERE id = $3
	`, invoice.Status,
		invoice.UpdatedAt,
		invoice.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrInvoiceNotFound
	}

	return nil
}
