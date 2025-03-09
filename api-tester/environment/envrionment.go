package environment

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

var environment string

func GetEnvironmentInfo() string {
	return environment
}

// COPY FUNCTION FROM PROJECTS PACKAGE TO AVOID IMPORT CYCLE
func getProjectsOptions() []huh.Option[string] {
	var projectOptions []huh.Option[string]

	folderPath := "data/projects"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Error reading directory", folderPath)
	}

	for _, file := range files {
		projectName := strings.TrimSuffix(file.Name(), ".json")
		projectOptions = append(projectOptions, huh.NewOption(projectName, projectName))
	}

	return projectOptions
}

func SetEnvironmentInfo() {
	projectOptions := getProjectsOptions()
	projectOptions = append(projectOptions, huh.NewOption("Unselect current environment", ""))

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select project to set as environment").
				Options(projectOptions...).
				Value(&environment),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	if environment != "" {
		fmt.Printf("\nProject %s selected as current environment", environment)
	} else {
		fmt.Println("Unselected current environment")
	}
	fmt.Fprintln(os.Stdout)
}
