package projects

import (
	"api-tester/api"
	"api-tester/environment"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
)

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
