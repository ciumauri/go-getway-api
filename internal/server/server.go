package server

import (
	"net/http"

	accountService "github.com/devfullcycle/imersao22/go-gateway/internal/account/service"
	accountHandlers "github.com/devfullcycle/imersao22/go-gateway/internal/account/web/handlers"
	"github.com/devfullcycle/imersao22/go-gateway/internal/middleware"

	invoiceService "github.com/devfullcycle/imersao22/go-gateway/internal/invoice/service"
	invoiceHandlers "github.com/devfullcycle/imersao22/go-gateway/internal/invoice/web/handlers"

	"github.com/go-chi/chi/v5"
)

// Server is a struct that represents a server
type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *accountService.AccountService
	invoiceService *invoiceService.InvoiceService
	port           string
}

// NewServer creates a new server
func NewServer(accountService *accountService.AccountService, invoiceService *invoiceService.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

// ConfigureRoutes configures the routes for the server
func (s *Server) ConfigureRoutes() {
	accountHandler := accountHandlers.NewAccountHandler(s.accountService)
	invoiceHandler := invoiceHandlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	// Account routes
	s.router.Post("/account", accountHandler.CreateAccount)
	s.router.Get("/account", accountHandler.GetAccount)

	// Invoice routes with authentication
	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		r.Post("/invoice", invoiceHandler.CreateInvoice)
		r.Get("/invoice/{id}", invoiceHandler.GetInvoiceByID)
		r.Get("/invoice", invoiceHandler.ListInvoicesByAccount)
	})
}

// Start starts the server
func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
