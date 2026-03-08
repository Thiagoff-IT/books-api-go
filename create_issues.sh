#!/bin/bash
set -e

echo "=== Criando issues ==="

gh issue create \
  --title "Adicionar validacao de campos obrigatorios na criacao/atualizacao de livros" \
  --label "enhancement" \
  --body "## Contexto
Atualmente, os endpoints POST e PUT /api/v1/livros aceitam payloads com campos vazios. E possivel criar um livro sem titulo, autor ou categoria.

## Proposta
- Validar que titulo, autor e categoria sao campos obrigatorios e nao vazios
- Retornar 400 Bad Request com mensagem descritiva dos campos faltantes
- Considerar limitar o tamanho dos campos (ex: titulo max 255 chars, sinopse max 2000 chars)

## Criterios de aceite
- [ ] Campos obrigatorios validados na criacao e na atualizacao
- [ ] Mensagens de erro claras indicando quais campos faltam
- [ ] Testes para cenarios de validacao"

echo "--- Issue 2 criada ---"

gh issue create \
  --title "Adicionar middleware de CORS" \
  --label "enhancement" \
  --body "## Contexto
A API nao possui configuracao de CORS, impedindo consumo por frontends em outros dominios.

## Proposta
- Adicionar middleware de CORS configuravel
- Permitir configurar origens permitidas via variavel de ambiente
- Incluir headers padrao (Content-Type, Authorization)

## Criterios de aceite
- [ ] Middleware de CORS aplicado a todas as rotas
- [ ] Origens configuraveis via env var
- [ ] Requisicoes OPTIONS (preflight) tratadas corretamente"

echo "--- Issue 3 criada ---"

gh issue create \
  --title "Implementar graceful shutdown do servidor HTTP" \
  --label "enhancement" \
  --body "## Contexto
O servidor usa log.Fatal(http.ListenAndServe(...)) sem suporte a graceful shutdown. Requisicoes em andamento podem ser interrompidas abruptamente ao encerrar o processo.

## Proposta
- Usar http.Server com contexto e signals (SIGINT, SIGTERM)
- Aguardar requisicoes em andamento antes de encerrar
- Timeout configuravel para shutdown

## Criterios de aceite
- [ ] Servidor captura sinais de encerramento
- [ ] Requisicoes ativas sao finalizadas antes do shutdown
- [ ] Log informativo sobre inicio e fim do shutdown"

echo "--- Issue 4 criada ---"

gh issue create \
  --title "Tornar porta do servidor configuravel via variavel de ambiente" \
  --label "enhancement" \
  --body "## Contexto
A porta esta hardcoded como :8000 em cmd/api/main.go. Isso dificulta deploy em ambientes com portas dinamicas.

## Proposta
- Ler porta da variavel de ambiente PORT (ou API_PORT)
- Usar :8000 como valor padrao caso a variavel nao esteja definida
- Considerar usar pacote de configuracao (envconfig ou similar)

## Criterios de aceite
- [ ] Porta configuravel via env var
- [ ] Valor padrao mantido quando variavel nao existe
- [ ] docker-compose.yml atualizado com a variavel"

echo "--- Issue 5 criada ---"

gh issue create \
  --title "Adicionar middleware de logging para requisicoes HTTP" \
  --label "enhancement" \
  --body "## Contexto
Nao ha logging de requisicoes HTTP. E dificil depurar problemas e monitorar o uso da API em producao.

## Proposta
- Adicionar middleware que loga: metodo, path, status code, duracao e IP
- Usar log estruturado (ex: slog do Go 1.21+)
- Logar em formato JSON para facilitar integracao com ferramentas de observabilidade

## Criterios de aceite
- [ ] Todas as requisicoes sao logadas
- [ ] Log inclui metodo, path, status, duracao
- [ ] Formato estruturado (JSON)"

echo "--- Issue 6 criada ---"

gh issue create \
  --title "Adicionar paginacao ao endpoint de listagem de livros" \
  --label "enhancement" \
  --body "## Contexto
O endpoint GET /api/v1/livros retorna todos os livros sem paginacao. Com o crescimento dos dados, isso pode causar problemas de performance e consumo de memoria.

## Proposta
- Suportar query params: page e limit (ex: ?page=1&limit=10)
- Retornar metadados de paginacao (total, page, limit, totalPages)
- Definir limite padrao (ex: 20) e maximo (ex: 100)

## Criterios de aceite
- [ ] Paginacao funcional com page e limit
- [ ] Metadados de paginacao na resposta
- [ ] Valores padrao sensatos quando params nao fornecidos"

echo "--- Issue 7 criada ---"

gh issue create \
  --title "Adicionar endpoint de health check" \
  --label "enhancement" \
  --body "## Contexto
Nao existe endpoint de health check. Isso dificulta monitoramento e uso com Docker/Kubernetes.

## Proposta
- Criar endpoint GET /health ou GET /api/v1/health
- Retornar status 200 com informacoes basicas (status, versao, uptime)
- Adicionar HEALTHCHECK no Dockerfile
- Util para readiness/liveness probes em K8s

## Criterios de aceite
- [ ] Endpoint /health retorna 200 OK
- [ ] Resposta inclui status e informacoes da aplicacao
- [ ] HEALTHCHECK configurado no Dockerfile"

echo "=== Todas as issues criadas com sucesso ==="
