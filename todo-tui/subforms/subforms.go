package subforms

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func ViewTodos(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No to dos added!")
	} else {
		fmt.Print("\n-----TODO LIST-----\n")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}

		fmt.Println()
	}
}

func AddTask() (string, error) {
	var task string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Add your task").
				Placeholder("eg. Add groceries").
				Value(&task),
		),
	)

	if err := form.Run(); err != nil {
		return "", err
	}

	return task, nil
}

func EditTask(tasks []string) ([]string, error) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to edit!")
		return tasks, nil
	}

	var taskIndex int
	options := make([]huh.Option[int], len(tasks))

	for i, task := range tasks {
		options[i] = huh.NewOption(fmt.Sprintf("%d %s", i+1, task), i)
	}

	if err := huh.NewSelect[int]().
		Title("Select a to do to edit").
		Options(options...).
		Value(&taskIndex).Run(); err != nil {
		return tasks, err
	}

	if err := huh.NewInput().
		Title("Edit todo").
		Value(&tasks[taskIndex]).Run(); err != nil {
		return tasks, err
	}

	fmt.Println("Task edited succesfully")
	return tasks, nil
}

func DeleteTask(tasks []string) ([]string, error) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to delete!")
		return tasks, nil
	}

	var taskIndex int
	var confirmed bool
	options := make([]huh.Option[int], len(tasks))

	for i, task := range tasks {
		options[i] = huh.NewOption(fmt.Sprintf("%d %s", i+1, task), i)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Select a to do to delete").
				Options(options...).
				Value(&taskIndex),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Are you sure you want to delete this todo?").
				Affirmative("Yes").
				Negative("No").
				Value(&confirmed),
		),
	)

	if err := form.Run(); err != nil {
		return tasks, err
	}

	if confirmed {
		tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
		fmt.Println("Task Deleted successfully")
	}

	return tasks, nil
}
