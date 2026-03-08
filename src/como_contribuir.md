# Como contribuir com novos exemplos

Este guia resume como criar ou atualizar tópicos sem quebrar o padrão do repositório.

## Fluxo de contribuição

1. Escolha um tópico pendente no `tasks.md`.
2. Leia `AGENTS.md` antes de começar.
3. Crie ou atualize um diretório de tópico com nome em inglês e minúsculo.
4. Escreva README em português do Brasil com comandos completos.
5. Adicione exemplos Go pequenos, executáveis e com tratamento de erro.
6. Inclua testes unitários quando fizer sentido.
7. Atualize navegação em `README.md` e `src/_indice.md`.
8. Execute as validações obrigatórias e só depois finalize.

## Convenções de nome

- Diretórios de tópico: `snake_case`.
- Arquivos de código: nomes curtos e descritivos em inglês.
- Identificadores Go: inglês e estilo idiomático.

## Reintegração de conteúdo antigo

Ao migrar algo de `nao_tratados/`:

- preserve a ideia didática principal do material original;
- modernize APIs obsoletas;
- mantenha exemplos mínimos e objetivos;
- remova o diretório antigo apenas após validar o novo tópico e atualizar a navegação.

## Checklist rápido antes do commit

- README claro e em português do Brasil.
- Código compilando e testes passando.
- Sem dependência externa sem justificativa.
- Links do índice e README principal atualizados.
