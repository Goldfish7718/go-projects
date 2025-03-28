package main

import (
	"fmt"
	"log"
	"todo-tui/subforms"

	"github.com/charmbracelet/huh"
)

var tasks []string

func main() {

	for {
		var choice string
		var err error

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("What would you like to do?").
					Options(
						huh.NewOption("View Todos", "view"),
						huh.NewOption("Add a task", "add"),
						huh.NewOption("Edit a task", "edit"),
						huh.NewOption("Delete a task", "delete"),
						huh.NewOption("Exit", "exit"),
					).
					Value(&choice),
			),
		)

		if err := form.Run(); err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "view":
			subforms.ViewTodos(tasks)

		case "add":
			task, err := subforms.AddTask()
			if err != nil {
				log.Println("Error adding task", err)
				continue
			}

			tasks = append(tasks, task)
			fmt.Println("Task added succesfully")

		case "edit":
			tasks, err = subforms.EditTask(tasks)

			if err != nil {
				log.Println("Error editing task", err)
			}

		case "delete":
			tasks, err = subforms.DeleteTask(tasks)

			if err != nil {
				log.Println("Error deleting task", err)
			}

		case "exit":
			fmt.Println("Goodbye!")
			return
		}
	}

}
