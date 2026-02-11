package main

import (
	"fmt"
	"log"
	"net/http"

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

	// Iniciar servidor
	porta := ":8000"
	fmt.Printf("Servidor rodando em http://localhost%s\n", porta)
	log.Fatal(http.ListenAndServe(porta, r))
}
