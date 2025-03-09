package projects

import (
	"api-tester/api"
	"api-tester/environment"
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

func SaveRequest(reqType string, route string) {
	var projectName string

	options := GetProjectsOptions()

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select project to save this request:").
				Options(options...).
				Value(&projectName),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	folderPath := "data/projects"
	filePath := filepath.Join(folderPath, projectName+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var project Project

	err = json.Unmarshal(data, &project)
	if err != nil {
		log.Fatal("Error Unmarshalling JSON", err)
	}

	newRequest := Request{
		ReqType: reqType,
		Route:   route,
	}

	project.Requests = append(project.Requests, newRequest)

	updatedData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling data", err)
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	fmt.Println("Request saved successfully")
}

func PerformSavedRequest() {
	projectName := environment.GetEnvironmentInfo()
	if projectName == "" {
		fmt.Println("No environment selected!")
		return
	}

	folderPath := "data/projects"
	filePath := filepath.Join(folderPath, projectName+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var project Project
	var requestOptions []huh.Option[int]
	var requestIndex int

	err = json.Unmarshal(data, &project)
	if err != nil {
		log.Fatal("Error Unmarshalling JSON", err)
	}

	for index, request := range project.Requests {
		requestOption := fmt.Sprintf("%s %s", request.ReqType, request.Route)
		requestOptions = append(requestOptions, huh.NewOption(requestOption, index))
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Select request to perform:\nProject: " + projectName + "\nBase URL: " + project.BaseUrl).
				Options(requestOptions...).
				Value(&requestIndex),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	selectedRequest := project.Requests[requestIndex]
	completeUrl := project.BaseUrl + selectedRequest.Route

	switch selectedRequest.ReqType {
	case "GET":
		api.Get(completeUrl)

	case "POST":
		api.Post(completeUrl)

	case "PUT":
		api.Put(completeUrl)

	case "DELETE":
		api.Delete(completeUrl)
	}
}
