package main

import (
	"log"

	"github.com/charmbracelet/huh"
)

func main() {
	var baseUrl string
	var choice string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Welcome to API tester. Enter your base URL:").
				Value(&baseUrl),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an action").
				Options(
					huh.NewOption("Make a new request", "new"),
					huh.NewOption("Exit", "exit"),
				).
				Value(&choice),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}
}
