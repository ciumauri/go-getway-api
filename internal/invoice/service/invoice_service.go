package service

import (
	"github.com/devfullcycle/imersao22/go-gateway/internal/account/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"github.com/devfullcycle/imersao22/go-gateway/internal/invoice/dto"
)

// InvoiceService is a service for invoices
type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    service.AccountService
}

// NewInvoiceService creates a new invoice service
func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService service.AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountService:    accountService,
	}
}

// CreateInvoice creates a new invoice
func (s *InvoiceService) CreateInvoice(input *dto.CreateInvoiceDTO) (*dto.InvoiceResponseDTO, error) {
	AccountResponseDTO, err := s.accountService.GetByApiKey(input.ApiKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoiceDTO(input, AccountResponseDTO.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == string(domain.StatusApproved) {
		_, err = s.accountService.UpdateBalance(input.ApiKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := s.invoiceRepository.CreateInvoice(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoiceDTO(invoice), nil
}

// GetInvoiceByID gets an invoice by ID
func (s *InvoiceService) GetByID(id, apiKey string) (*dto.InvoiceResponseDTO, error) {
	invoice, err := s.invoiceRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	accountResponseDTO, err := s.accountService.GetByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountResponseDTO.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	return dto.FromInvoiceDTO(invoice), nil
}

// ListByAccountID lists all invoices by account ID
func (s *InvoiceService) ListByAccountID(accountID string) ([]*dto.InvoiceResponseDTO, error) {
	invoices, err := s.invoiceRepository.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceResponseDTO, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoiceDTO(invoice)
	}

	return output, nil
}

// ListByAccountAPIKey lists all invoices by account API key
func (s *InvoiceService) ListByAccountAPIKey(apiKey string) ([]*dto.InvoiceResponseDTO, error) {
	accountResponseDTO, err := s.accountService.GetByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccountID(accountResponseDTO.ID)
}
