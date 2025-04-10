package middleware

import (
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/account/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

// AuthMiddleware is a middleware that checks if the request has a valid API key
type AuthMiddleware struct {
	accountService service.AccountService
}

// NewAuthMiddleware creates a new AuthMiddleware
func NewAuthMiddleware(accountService service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{
		accountService: accountService,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "X-API-Key header is required", http.StatusUnauthorized)
			return
		}

		_, err := m.accountService.GetByApiKey(apiKey)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
