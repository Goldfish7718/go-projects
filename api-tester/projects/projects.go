package projects

import (
	"api-tester/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
)

type Project struct {
	ProjectName string    `json:"projectName"`
	BaseUrl     string    `json:"baseUrl"`
	Requests    []Request `json:"requests"`
}

type Request struct {
	ReqType string `json:"type"`
	Route   string `json:"route"`
}

func New() {
	var projectName string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter Project name:").
				Value(&projectName),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	baseUrl := utils.GetBaseUrl()

	project := Project{
		ProjectName: projectName,
		BaseUrl:     baseUrl,
		Requests:    []Request{},
	}

	folderPath := "data/projects"

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating directory", err)
	}

	filePath := filepath.Join(folderPath, projectName+".json")
	fileData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		log.Fatal("❌ Error marshalling JSON:", err)
	}

	err = os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	fmt.Printf("Succefully created project %s\n", projectName)
}

func View() {
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

	for index, projectName := range projects {
		fmt.Printf("\n%d. %s\n", index+1, projectName)
	}
}

func Delete() {
	var projectToDelete string
	var projectDeleteConfirm bool
	projectOptions := GetProjectsOptions()

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select project to delete:").
				Options(projectOptions...).
				Value(&projectToDelete),

			huh.NewConfirm().
				Title("Are you sure you want to delete this project?").
				Affirmative("Yes").
				Negative("No").
				Value(&projectDeleteConfirm),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	if !projectDeleteConfirm {
		return
	}

	folderPath := "data/projects"
	filepath := filepath.Join(folderPath, projectToDelete+".json")

	err := os.Remove(filepath)
	if err != nil {
		log.Fatal("Error deleting file", err)
	}

	fmt.Printf("\nProject %s deleted succesfully", projectToDelete)
}
