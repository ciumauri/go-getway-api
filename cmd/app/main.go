package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	accountRepository "github.com/devfullcycle/imersao22/go-gateway/internal/account/repository"
	accountService "github.com/devfullcycle/imersao22/go-gateway/internal/account/service"

	invoiceRepository "github.com/devfullcycle/imersao22/go-gateway/internal/invoice/repository"
	invoiceService "github.com/devfullcycle/imersao22/go-gateway/internal/invoice/service"

	"github.com/devfullcycle/imersao22/go-gateway/internal/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// carrega variável de ambiente com fallback
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// cria conexão com banco de dados
func initDB() *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Erro ao abrir conexão com banco: %v", err)
	}

	// Testar conexão
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco: %v", err)
	}

	return db
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Arquivo .env não encontrado. Usando variáveis de ambiente padrão.")
	}

	db := initDB()
	defer db.Close()

	// Inicializar repos e services
	accRepo := accountRepository.NewAccountRepository(db)
	accService := accountService.NewAccountService(accRepo)

	invRepo := invoiceRepository.NewInvoiceRepository(db)
	invService := invoiceService.NewInvoiceService(invRepo, accService)

	// Inicializar servidor
	port := getEnv("PORT", "8080")
	srv := server.NewServer(accService, invService, port)

	srv.ConfigureRoutes()

	log.Printf("🚀 Servidor iniciado na porta %s", port)
	if err := srv.Start(); err != nil {
		log.Fatalf("❌ Erro ao iniciar servidor: %v", err)
	}
}
