# Task CLI – Gerenciador de Tarefas em Linha de Comando

Um pequeno aplicativo Go que permite criar, listar, marcar como concluída e deletar tarefas totalmente usando apenas a biblioteca padrão (stdlib). O projeto não tem dependências externas.

## Estrutura do projeto

- `main.go` – ponto de entrada da aplicação.
- `go.mod` – módulo Go definido como `task-cli`.
- `tasks.json` – arquivo de persistência (gerado e carregado automaticamente).
- `.gitignore` – ignora binários compilados e o arquivo `tasks.json`.
- `README.md` – este arquivo.

## Compilar

```bash
go build -o task-cli main.go
```

O executável gerado pode ser executado diretamente:

```bash
./task-cli
```

Ou com argumentos CLI:

```bash
./task-cli add "Lembrar comprar leite"
./task-cli list
./task-cli done 1
./task-cli delete 2
./task-cli help
```

## Comandos

| Comando | Descrição |
|---|---|
| **`add "<descrição>"`** | Adiciona uma nova tarefa com ID incremental, status pendente e data de criação. |
| **`list`** | Lista todas as tarefas mostrando ID, status ([ ] ou [x]) e descrição. |
| **`done <id>`** | Marca a tarefa com o ID informado como concluída. |
| **`delete <id>`** | Remove a tarefa com o ID informado. |
| **`help`** | Mostra as instruções de uso. |

### Exemplos completos

```bash
$ ./task-cli add "Pagar contas de luz"
Tarefa 1 adicionada com status pendente.

$ ./task-cli list
--- Lista de tarefas ---
1 | [ ] | Pagar contas de luz

$ ./task-cli done 1
Tarefa 1 marcada como concluída.

$ ./task-cli list
1 | [x] | Pagar contas de luz

$ ./task-cli delete 1
Tarefa 1 removida.

$ ./task-cli help
Uso do task-cli ...
```

## Persistência

As tarefas são salvas e carregadas automaticamente no arquivo `tasks.json` usando `encoding/json` e `os`. Cada tarefa guarda seu ID, status e descrição, além da data de criação em UTC.
