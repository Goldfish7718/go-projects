package subforms

import (
	"api-tester/projects"
	"api-tester/utils"
	"log"

	"github.com/charmbracelet/huh"
)

func RequestSubform() {
	var requestType string
	var route string
	var save bool

	baseUrl := utils.GetBaseUrl()

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

	case "POST":
		utils.Post(completeUrl)

	case "PUT":
		utils.Put(completeUrl)

	case "DELETE":
		utils.Delete(completeUrl)
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to save this request?").
				Affirmative("Yes").
				Negative("No").
				Value(&save),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	if save {
		projects.SaveRequest(requestType, route)
	}
}

func ProjectSubform() {
	var action string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Request type").
				Options(
					huh.NewOption("Create new Project", "new"),
					huh.NewOption("View Projects", "view"),
				).
				Value(&action),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	switch action {
	case "new":
		projects.New()

	case "view":
		projects.View()
	}
}
