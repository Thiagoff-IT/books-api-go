package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IamThiago-IT/Go/internal/handler"
	"github.com/IamThiago-IT/Go/internal/repository"
	"github.com/IamThiago-IT/Go/internal/router"
	"github.com/IamThiago-IT/Go/internal/service"
)

func main() {
	// Inicializar camadas (injeção de dependência manual)
	livroRepo := repository.NovoLivroRepository()
	livroService := service.NovoLivroService(livroRepo)
	livroHandler := handler.NovoLivroHandler(livroService)

	// Configurar router
	r := router.NovoRouter(livroHandler)

	// Porta configurável via variável de ambiente
	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "8000"
	}

	srv := &http.Server{
		Addr:         ":" + porta,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Iniciar servidor em goroutine
	go func() {
		slog.Info("servidor iniciado", "porta", porta)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("erro ao iniciar servidor", "erro", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("encerrando servidor...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("erro no shutdown", "erro", err)
		os.Exit(1)
	}
	slog.Info("servidor encerrado")
}
