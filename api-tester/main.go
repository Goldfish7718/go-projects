package main

import (
	"api-tester/subforms"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func getBaseUrl() string {
	var baseUrl string = "http://localhost:3000"

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter your base URL:").
				Value(&baseUrl),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	return baseUrl
}

func main() {
	var choice string

	fmt.Println("Welcome to API tester!")

	baseUrl := getBaseUrl()

	for {
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select an action").
					Options(
						huh.NewOption("Make a new request", "new"),
						huh.NewOption("Set base url", "baseUrl"),
						huh.NewOption("Exit", "exit"),
					).
					Value(&choice),
			),
		).Run(); err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "new":
			subforms.RequestSubform(baseUrl)

		case "baseUrl":
			baseUrl = getBaseUrl()

		case "exit":
			return
		}
	}
}
