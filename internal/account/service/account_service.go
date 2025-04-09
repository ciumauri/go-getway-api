package service

import (
	"github.com/devfullcycle/imersao22/go-gateway/internal/account/dto"
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

type AccountService struct {
	accountRepository domain.AccountRepository
}

func NewAccountService(accountRepository domain.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountDTO) (*dto.AccountResponseDTO, error) {
	account := dto.ToAccountDTO(input)

	existingAccount, err := s.accountRepository.GetByApiKey(account.ApiKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}

	err = s.accountRepository.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccountDTO(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountResponseDTO, error) {
	account, err := s.accountRepository.GetByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.accountRepository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccountDTO(account)
	return &output, nil
}

func (s *AccountService) GetByApiKey(apiKey string) (*dto.AccountResponseDTO, error) {
	account, err := s.accountRepository.GetByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccountDTO(account)
	return &output, nil
}

func (s *AccountService) GetByID(id string) (*dto.AccountResponseDTO, error) {
	account, err := s.accountRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccountDTO(account)
	return &output, nil
}
