package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/chzyer/readline"
)

func AcceptRequestBody() *bytes.Buffer {

	var jsonStr string

	rl, err := readline.New("> ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer rl.Close()

	fmt.Println("Type your JSON (press Enter to submit):")

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		if line == "" {
			break
		}
		jsonStr += line
	}

	fmt.Println("You entered:")
	fmt.Println(jsonStr)

	reqBody := bytes.NewBuffer([]byte(jsonStr))
	return reqBody
}

func ParseResponseBody(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	return string(body)
}

func GetBaseUrl() string {
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
