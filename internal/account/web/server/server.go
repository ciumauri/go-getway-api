package server

import (
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/account/service"
	"github.com/devfullcycle/imersao22/go-gateway/internal/account/web/handlers"
	"github.com/go-chi/chi/v5"
)

// Server is a struct that represents a server
type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

// NewServer creates a new server
func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
}

// ConfigureRoutes configures the routes for the server
func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", accountHandler.CreateAccount)
	s.router.Get("/accounts", accountHandler.GetAccount)
}

// Start starts the server
func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
