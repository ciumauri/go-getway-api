package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/account/dto"
	"github.com/devfullcycle/imersao22/go-gateway/internal/account/service"
)

// AccountHandler is a struct that represents an account handler
type AccountHandler struct {
	accountService *service.AccountService
}

// NewAccountHandler creates a new account handler
func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

// CreateAccount creates a new account
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateAccountDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.accountService.CreateAccount(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// GetAccount gets an account by its API key
func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "API key is required", http.StatusUnauthorized)
		return
	}

	output, err := h.accountService.GetByApiKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
