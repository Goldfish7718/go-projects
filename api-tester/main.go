package main

import (
	"api-tester/environment"
	"api-tester/projects"
	"api-tester/subforms"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func main() {
	var choice string
	var selectActionString string

	fmt.Println("Welcome to API tester!")

	for {
		environmentInfo := environment.GetEnvironmentInfo()
		if environmentInfo != "" {
			selectActionString = fmt.Sprintf("(%s) Select an action:", environmentInfo)
		} else {
			selectActionString = "Select an action:"
		}

		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title(selectActionString).
					Options(
						huh.NewOption("Make a new request", "new_request"),
						huh.NewOption("Set Environment", "set_env"),
						huh.NewOption("Perform saved request", "perform_saved"),
						huh.NewOption("Exit", "exit"),
					).
					Value(&choice),
			),
		).Run(); err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "new_request":
			subforms.RequestSubform()

		case "set_env":
			environment.SetEnvironmentInfo()

		case "perform_saved":
			projects.PerformSavedRequest()

		case "exit":
			fmt.Println("Bye!")
			return
		}
	}
}
