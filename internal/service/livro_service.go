package service

import (
	"github.com/IamThiago-IT/Go/internal/model"
	"github.com/IamThiago-IT/Go/internal/repository"
)

// LivroService define o contrato das regras de negócio para livros.
type LivroService interface {
	CriarLivro(livro model.Livro) model.Livro
	BuscarLivroPorID(id int) (model.Livro, error)
	ListarLivros() []model.Livro
	AtualizarLivro(id int, livro model.Livro) (model.Livro, error)
	DeletarLivro(id int) error
}

type livroService struct {
	repo repository.LivroRepository
}

// NovoLivroService cria uma nova instância do serviço de livros.
func NovoLivroService(repo repository.LivroRepository) LivroService {
	return &livroService{repo: repo}
}

func (s *livroService) CriarLivro(livro model.Livro) model.Livro {
	return s.repo.Criar(livro)
}

func (s *livroService) BuscarLivroPorID(id int) (model.Livro, error) {
	return s.repo.BuscarPorID(id)
}

func (s *livroService) ListarLivros() []model.Livro {
	return s.repo.BuscarTodos()
}

func (s *livroService) AtualizarLivro(id int, livro model.Livro) (model.Livro, error) {
	return s.repo.Atualizar(id, livro)
}

func (s *livroService) DeletarLivro(id int) error {
	return s.repo.Deletar(id)
}
