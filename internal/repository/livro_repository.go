package repository

import (
	"errors"
	"sync"

	"github.com/IamThiago-IT/Go/internal/model"
)

var (
	ErrLivroNaoEncontrado = errors.New("livro não encontrado")
)

// LivroRepository define o contrato de acesso a dados para livros.
type LivroRepository interface {
	Criar(livro model.Livro) model.Livro
	BuscarPorID(id int) (model.Livro, error)
	BuscarTodos() []model.Livro
	Atualizar(id int, livro model.Livro) (model.Livro, error)
	Deletar(id int) error
}

// livroRepositoryMemoria é a implementação in-memory do repositório.
type livroRepositoryMemoria struct {
	mu      sync.RWMutex
	livros  []model.Livro
	nextID  int
}

// NovoLivroRepository cria uma nova instância do repositório em memória.
func NovoLivroRepository() LivroRepository {
	return &livroRepositoryMemoria{
		livros: make([]model.Livro, 0),
		nextID: 1,
	}
}

func (r *livroRepositoryMemoria) Criar(livro model.Livro) model.Livro {
	r.mu.Lock()
	defer r.mu.Unlock()

	livro.ID = r.nextID
	r.nextID++
	r.livros = append(r.livros, livro)
	return livro
}

func (r *livroRepositoryMemoria) BuscarPorID(id int) (model.Livro, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, livro := range r.livros {
		if livro.ID == id {
			return livro, nil
		}
	}
	return model.Livro{}, ErrLivroNaoEncontrado
}

func (r *livroRepositoryMemoria) BuscarTodos() []model.Livro {
	r.mu.RLock()
	defer r.mu.RUnlock()

	resultado := make([]model.Livro, len(r.livros))
	copy(resultado, r.livros)
	return resultado
}

func (r *livroRepositoryMemoria) Atualizar(id int, livro model.Livro) (model.Livro, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, item := range r.livros {
		if item.ID == id {
			livro.ID = id
			r.livros[i] = livro
			return livro, nil
		}
	}
	return model.Livro{}, ErrLivroNaoEncontrado
}

func (r *livroRepositoryMemoria) Deletar(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, item := range r.livros {
		if item.ID == id {
			r.livros = append(r.livros[:i], r.livros[i+1:]...)
			return nil
		}
	}
	return ErrLivroNaoEncontrado
}
