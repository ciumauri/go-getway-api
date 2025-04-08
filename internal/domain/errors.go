package domain

import "errors"

var (
	// ErrAccountNotFound is returned when an account is not found
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicateAccount is returned when a duplicate account is found
	ErrDuplicateAccount = errors.New("account already exists")
	// ErrDuplicatedAPIKey is returned when a duplicated API key is found
	ErrDuplicatedAPIKey = errors.New("API key already exists")
	// ErrInvoiceNotFound is returned when an invoice is not found
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAccess is returned when a user is not authorized to access a resource
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
