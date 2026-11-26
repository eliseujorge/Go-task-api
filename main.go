package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Tarefa representa uma única entrada no gerenciador de tarefas.
type Tarefa struct {
	ID        string    `json:"id"`
	Status   string    `json:"status"` // "pendente" ou "concluída"
	Descrição string   `json:"descrição"`
	DataCriacao time.Time `json:"data_criacao"`
}

// loadTasks lê o arquivo tasks.json (se existir) e devolve uma fatia de Tarefa.
func loadTasks() ([]Tarefa, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Tarefa{}, nil
		}
		return nil, err
	}
	var tasks []Tarefa
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// saveTasks escreve a fatia de Tarefa de volta em tasks.json.
func saveTasks(tasks []Tarefa) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

// printList exibe todas as tarefas na forma ID – status – descrição.
func printList(tasks []Tarefa) {
	fmt.Println("\n--- Lista de tarefas ---")
	for _, t := range tasks {
		var statusStr string
		if t.Status == "pendente" {
			statusStr = "[ ]"
		} else {
			statusStr = "[x]"
		}
		fmt.Printf("%s | %s | %s\n", t.ID, statusStr, t.Descrição)
	}
}

// main executa o CLI. Os argumentos passados são os elementos de os.Args[1:].
func main() {
	args := os.Args[1:]

	// Carrega as tarefas existentes.
	tasks, err := loadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao carregar tarefas: %v\n", err)
		return
	}

	// Exibe a lista atual.
	printList(tasks)

	// Processa os comandos.
	if len(args) == 0 {
		fmt.Println("Comando desconhecido: vazio")
		fmt.Println("Use 'help' para ver as opções disponíveis.")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Uso: task-cli add \"descrição da tarefa\"\n")
			return
		}
		desc := args[1]
		newID := strconv.Itoa(len(tasks) + 1)
		tasks = append(tasks, Tarefa{
			ID:        newID,
			Status:   "pendente",
			Descrição: desc,
			DataCriacao: time.Now().UTC(),
		})
		saveTasks(tasks)
		fmt.Printf("Tarefa %s adicionada com status pendente.\n", newID)

	case "list":
		printList(tasks)

	case "done":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Uso: task-cli done <id>\n")
			return
		}
		id := args[1]
		found := false
		for i, t := range tasks {
			if t.ID == id {
				t.Status = "concluída"
				found = true
				break
			}
		}
		if found {
			saveTasks(tasks)
			fmt.Printf("Tarefa %s marcada como concluída.\n", id)
		} else {
			fmt.Fprintf(os.Stderr, "Erro: nenhuma tarefa com ID '%s' encontrada.\n", id)
		}

	case "delete":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Uso: task-cli delete <id>\n")
			return
		}
		id := args[1]
		found := false
		for i, t := range tasks {
			if t.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				found = true
				break
			}
		}
		if found {
			saveTasks(tasks)
			fmt.Printf("Tarefa %s removida.\n", id)
		} else {
			fmt.Fprintf(os.Stderr, "Erro: nenhuma tarefa com ID '%s' encontrada.\n", id)
		}

	case "help":
		fmt.Println("Uso do task-cli (CLI para gerenciar tarefas):")
		fmt.Println("  add \"descrição da tarefa\"   - adiciona uma nova tarefa com ID incremental, status pendente e data de criação.")
		fmt.Println("  list                          - lista todas as tarefas mostrando ID, status ([ ] ou [x]) e descrição.")
		fmt.Println("  done <id>                     - marca a tarefa com o ID informado como concluída.")
		fmt.Println("  delete <id>                   - remove a tarefa com o ID informado.")
		fmt.Println("  help                          - mostra estas instruções.")
	default:
		fmt.Fprintf(os.Stderr, "Comando desconhecido: %s\n", args[0])
		fmt.Println("Use 'help' para ver as opções disponíveis.")
	}
}
