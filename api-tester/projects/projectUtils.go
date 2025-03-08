package projects

import (
	"log"
	"os"
	"strings"
)

func GetProjects() []string {
	var projects []string

	folderPath := "data/projects"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Error reading directory", folderPath)
	}

	for _, file := range files {
		projectName := strings.TrimSuffix(file.Name(), ".json")
		projects = append(projects, projectName)
	}

	return projects
}
