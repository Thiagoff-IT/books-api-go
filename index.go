package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Livro struct {
	ID       int    `json:"id"`
	Titulo   string `json:"titulo"`
	Categoria string `json:"categoria"`
	Autor    string `json:"autor"`
	Sinopse  string `json:"sinopse"`
}

var livros []Livro
var idAtual int

func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/livros", criarLivro).Methods("POST")
	router.HandleFunc("/livros/{id}", atualizarLivro).Methods("PUT")
	router.HandleFunc("/livros/{id}", listarLivro).Methods("GET")
	router.HandleFunc("/livros", listarTodosLivros).Methods("GET")
	router.HandleFunc("/livros/{id}", deletarLivro).Methods("DELETE")
	
	http.ListenAndServe(":8000", router)
}

func criarLivro(w http.ResponseWriter, r *http.Request) {
	var livro Livro
	json.NewDecoder(r.Body).Decode(&livro)
	idAtual++
	livro.ID = idAtual
	livros = append(livros, livro)
	json.NewEncoder(w).Encode(livro)
}

func atualizarLivro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range livros {
		if item.ID == id {
			livros = append(livros[:index], livros[index+1:]...)
			var livro Livro
			json.NewDecoder(r.Body).Decode(&livro)
			livro.ID = id
			livros = append(livros, livro)
			json.NewEncoder(w).Encode(livro)
			return
		}
	}
	json.NewEncoder(w).Encode(&Livro{})
}

func listarLivro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range livros {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Livro{})
}

func listarTodosLivros(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(livros)
}

func deletarLivro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range livros {
		if item.ID == id {
			livros = append(livros[:index], livros[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(livros)
}

