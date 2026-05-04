package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"encoding/json"
	"todo-app/todo"
)

const todosFilePath = "todo/todos.json"

type todosType struct {
	Todos []todo.Todo
}

func main() {

	for {
		
		buf := bufio.NewReader(os.Stdin)

		fmt.Println("Enter a command (add, list, quit):")

		command, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command:", err)
			return
		}

		command = strings.TrimSpace(command)

		commandWords := strings.Fields(command)

		if commandWords[0] == "quit" {
			fmt.Println("Quitting...")
			return
		}

		todos, err := os.ReadFile(todosFilePath)
		if err != nil {
			fmt.Println("Error reading todos file:", err)
			return
		}

		var savedTodos todosType
		if err := json.Unmarshal(todos, &savedTodos); err != nil {
			fmt.Println("Error unmarshalling savedtodos:", err)
			return
		}

		if commandWords[0] == "add" {
			// newTodo := commandWords[1]

			todo := todo.Todo {
				Text: strings.Join(commandWords[1:], " ") ,
				Done: false,
			}

			fmt.Println("saved todos: ", savedTodos)

			savedTodos.Todos = append(savedTodos.Todos, todo)

			newTodos, err := json.Marshal(savedTodos)

			if err != nil {
				fmt.Println("Error marshalling new todos:", err)
				return
			}

			if err := os.WriteFile(todosFilePath, newTodos, 0644); err != nil {
				fmt.Println("Error writing todos file:", err)
				return
			}

		} else if commandWords[0] == "list" {
			fmt.Println("Listing todos:")
			for idx, todo := range savedTodos.Todos {
				fmt.Println(idx+1, ": ", todo.Text)
			}
		} else {
			fmt.Println("Unknown command:", commandWords[0])
		}
	}
}
