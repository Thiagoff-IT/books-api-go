# Api Go

Problema a ser resolvido com este projeto
- desenvolva uma API Rest que contemple os livros de um usuário.

Requisitos:

- Como usuário eu gostaria de cadastrar um livro.
- Como usuário eu gostaria de atualizar um livro.
- Como usuário eu gostaria de listar um livro.
- Como usuário eu gostaria de listar todos os livros.
- Como usuário eu gostaria de deletar um livro.

Os livros devem ter:
```
- Título
- Categoria.
- Autor.
- Sinopse.
```

## Instalação e Execução

1. Clone o repositório:
    ```sh
    git clone https://github.com/seu-usuario/seu-repositorio.git
    ```
2. Navegue até o diretório do projeto:
    ```sh
    cd seu-repositorio
    ```
3. Instale as dependências:
    ```sh
    go mod tidy
    ```
4. Execute a aplicação:
    ```sh
    go run main.go
    ```

## Exemplos de Uso

### Cadastrar um Livro

```sh
curl -X POST http://localhost:8080/livros -H "Content-Type: application/json" -d '{
    "titulo": "O Senhor dos Anéis",
    "categoria": "Fantasia",
    "autor": "J.R.R. Tolkien",
    "sinopse": "Uma aventura épica na Terra Média."
}'
```

### Atualizar um Livro

```sh
curl -X PUT http://localhost:8080/livros/1 -H "Content-Type: application/json" -d '{
    "titulo": "O Hobbit",
    "categoria": "Fantasia",
    "autor": "J.R.R. Tolkien",
    "sinopse": "A aventura de Bilbo Bolseiro."
}'
```

### Listar um Livro

```sh
curl -X GET http://localhost:8080/livros/1
```

### Listar Todos os Livros

```sh
curl -X GET http://localhost:8080/livros
```

### Deletar um Livro

```sh
curl -X DELETE http://localhost:8080/livros/1
```
