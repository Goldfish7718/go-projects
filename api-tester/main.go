package main

import (
	"api-tester/subforms"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func main() {
	var choice string

	fmt.Println("Welcome to API tester!")

	for {
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select an action").
					Options(
						huh.NewOption("Make a new request", "newRequest"),
						huh.NewOption("Manage Projects", "manageProject"),
						huh.NewOption("Exit", "exit"),
					).
					Value(&choice),
			),
		).Run(); err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "newRequest":
			subforms.RequestSubform()

		case "manageProject":
			subforms.ProjectSubform()

		case "exit":
			fmt.Println("Bye!")
			return
		}
	}
}
