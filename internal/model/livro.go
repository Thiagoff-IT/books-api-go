package model

// Livro representa a entidade de domínio de um livro.
type Livro struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Categoria string `json:"categoria"`
	Autor     string `json:"autor"`
	Sinopse   string `json:"sinopse"`
}
