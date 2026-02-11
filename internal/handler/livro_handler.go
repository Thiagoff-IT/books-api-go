package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/IamThiago-IT/Go/internal/model"
	"github.com/IamThiago-IT/Go/internal/repository"
	"github.com/IamThiago-IT/Go/internal/service"
	"github.com/IamThiago-IT/Go/pkg/response"
)

// LivroHandler contém os handlers HTTP para o recurso livro.
type LivroHandler struct {
	service service.LivroService
}

// NovoLivroHandler cria uma nova instância do handler de livros.
func NovoLivroHandler(svc service.LivroService) *LivroHandler {
	return &LivroHandler{service: svc}
}

// CriarLivro godoc — POST /livros
func (h *LivroHandler) CriarLivro(w http.ResponseWriter, r *http.Request) {
	var livro model.Livro
	if err := json.NewDecoder(r.Body).Decode(&livro); err != nil {
		response.Erro(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	criado := h.service.CriarLivro(livro)
	response.JSON(w, http.StatusCreated, criado)
}

// ListarTodosLivros godoc — GET /livros
func (h *LivroHandler) ListarTodosLivros(w http.ResponseWriter, r *http.Request) {
	livros := h.service.ListarLivros()
	response.JSON(w, http.StatusOK, livros)
}

// ListarLivro godoc — GET /livros/{id}
func (h *LivroHandler) ListarLivro(w http.ResponseWriter, r *http.Request) {
	id, err := h.extrairID(r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, "ID inválido")
		return
	}

	livro, err := h.service.BuscarLivroPorID(id)
	if err != nil {
		if err == repository.ErrLivroNaoEncontrado {
			response.Erro(w, http.StatusNotFound, "livro não encontrado")
			return
		}
		response.Erro(w, http.StatusInternalServerError, "erro interno")
		return
	}

	response.JSON(w, http.StatusOK, livro)
}

// AtualizarLivro godoc — PUT /livros/{id}
func (h *LivroHandler) AtualizarLivro(w http.ResponseWriter, r *http.Request) {
	id, err := h.extrairID(r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var livro model.Livro
	if err := json.NewDecoder(r.Body).Decode(&livro); err != nil {
		response.Erro(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	atualizado, err := h.service.AtualizarLivro(id, livro)
	if err != nil {
		if err == repository.ErrLivroNaoEncontrado {
			response.Erro(w, http.StatusNotFound, "livro não encontrado")
			return
		}
		response.Erro(w, http.StatusInternalServerError, "erro interno")
		return
	}

	response.JSON(w, http.StatusOK, atualizado)
}

// DeletarLivro godoc — DELETE /livros/{id}
func (h *LivroHandler) DeletarLivro(w http.ResponseWriter, r *http.Request) {
	id, err := h.extrairID(r)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.service.DeletarLivro(id); err != nil {
		if err == repository.ErrLivroNaoEncontrado {
			response.Erro(w, http.StatusNotFound, "livro não encontrado")
			return
		}
		response.Erro(w, http.StatusInternalServerError, "erro interno")
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// extrairID extrai e converte o parâmetro {id} da rota.
func (h *LivroHandler) extrairID(r *http.Request) (int, error) {
	params := mux.Vars(r)
	return strconv.Atoi(params["id"])
}
