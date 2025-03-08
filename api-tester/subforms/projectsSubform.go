package subforms

import (
	"api-tester/projects"
	"log"

	"github.com/charmbracelet/huh"
)

func ProjectSubform() {
	var action string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Request type").
				Options(
					huh.NewOption("Create new Project", "new"),
					huh.NewOption("View Projects", "view"),
					huh.NewOption("Perform saved request", "req"),
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

	case "req":
		projects.PerformSavedRequest()
	}
}
