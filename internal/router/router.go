package router

import (
	"github.com/gorilla/mux"

	"github.com/IamThiago-IT/Go/internal/handler"
)

// NovoRouter configura e retorna o router com todas as rotas registradas.
func NovoRouter(livroHandler *handler.LivroHandler) *mux.Router {
	r := mux.NewRouter()

	// Prefixo da API
	api := r.PathPrefix("/api/v1").Subrouter()

	// Rotas de livros
	api.HandleFunc("/livros", livroHandler.CriarLivro).Methods("POST")
	api.HandleFunc("/livros", livroHandler.ListarTodosLivros).Methods("GET")
	api.HandleFunc("/livros/{id}", livroHandler.ListarLivro).Methods("GET")
	api.HandleFunc("/livros/{id}", livroHandler.AtualizarLivro).Methods("PUT")
	api.HandleFunc("/livros/{id}", livroHandler.DeletarLivro).Methods("DELETE")

	return r
}
