package subforms

import (
	"api-tester/utils"
	"log"

	"github.com/charmbracelet/huh"
)

func RequestSubform(baseUrl string) {
	var requestType string
	var route string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Request type").
				Options(
					huh.NewOption("GET", "GET"),
					huh.NewOption("POST", "POST"),
					huh.NewOption("PUT", "PUT"),
					huh.NewOption("DELETE", "DELETE"),
				).
				Value(&requestType),

			huh.NewInput().
				Title("Enter route (with trailing slash; appended to base URL)").
				Value(&route),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	completeUrl := baseUrl + route

	switch requestType {
	case "GET":
		utils.Get(completeUrl)

	case "DELETE":
		utils.Delete(completeUrl)

	case "POST":
		utils.Post(completeUrl)
	}
}
