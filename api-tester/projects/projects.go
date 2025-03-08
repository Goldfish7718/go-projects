package projects

import (
	"api-tester/api"
	"api-tester/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	projects := GetProjects()

	for index, projectName := range projects {
		fmt.Printf("\n%d. %s\n", index+1, projectName)
	}
}

func SaveRequest(reqType string, route string) {
	var projectName string
	var options []huh.Option[string]

	projects := GetProjects()

	for _, pName := range projects {
		options = append(options, huh.NewOption(pName, pName))
	}

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
	var projectName string
	var options []huh.Option[string]

	projects := GetProjects()

	for _, pName := range projects {
		options = append(options, huh.NewOption(pName, pName))
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select project:").
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
