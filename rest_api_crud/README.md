# API RESTful

Este tópico mostra um CRUD HTTP pequeno usando só `net/http`, `encoding/json` e um mapa protegido por `sync.Mutex`. Não tem banco, roteador externo nem middleware: a ideia é enxergar o contrato HTTP sem esconder o básico atrás de framework.

O exemplo trabalha com `/items` e `/items/{id}`:

- `POST /items` cria um item com `name` obrigatório;
- `GET /items` lista os itens em ordem de `id`, para o teste não depender da ordem interna do mapa;
- `GET /items/{id}` lê um item;
- `PUT /items/{id}` troca o nome mantendo o mesmo `id`;
- `DELETE /items/{id}` remove o item e responde `204 No Content`.

## Pré-requisitos

Go 1.26 ou mais novo.

## Rodar o exemplo

```bash
go run .
```

Saída esperada:

```text
POST /items -> 201
GET /items/1 -> 200
PUT /items/1 -> 200
DELETE /items/1 -> 204
```

O `main` não sobe uma porta TCP. Ele usa `httptest` para chamar o handler em memória e terminar rápido, o que é melhor para estudo e para CI. Se quiser transformar isso em servidor real depois, o ponto de troca é pequeno: `http.ListenAndServe` recebendo `store.Handler`.

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

O teste cobre o ciclo criar, ler, atualizar, listar e apagar. Também trava um caso chato de API pequena: payload com nome vazio precisa falhar com `400 Bad Request`, não criar lixo no mapa.

## Pontos de atenção

`map` em Go não preserva ordem. Por isso a listagem ordena por `id` antes de gerar JSON. Sem essa linha, o exemplo pode passar hoje e falhar quando tiver mais dados.

O armazenamento é em memória. Reiniciou o processo, perdeu tudo. Isso é proposital: persistência mudaria o assunto para SQL, arquivo ou transação, e este tópico é sobre método HTTP, status code e JSON.

Também não há autenticação, paginação, validação rica ou log estruturado. Tudo isso existe em APIs reais, mas aqui seria barulho antes da hora.

## Próximos passos

- Trocar o mapa por `database/sql` em outro tópico.
- Separar o handler em pacote próprio quando houver mais rotas.
- Adicionar `PATCH` só quando fizer sentido discutir atualização parcial.
