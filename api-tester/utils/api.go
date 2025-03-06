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
	"github.com/tidwall/pretty"
)

func Get(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	fmt.Print("Content type:", contentType, "\n")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	switch contentType {
	case "text/html; charset=utf-8":
		fmt.Println("Response:\n", string(body))

	case "application/json; charset=utf-8":
		prettyJson := pretty.Pretty(body)
		fmt.Println("Response:\n", string(prettyJson))
	}
}

func Post(url string) {
	var addBody string
	var jsonStr string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Add request body?").
				Options(
					huh.NewOption("Yes", "y"),
					huh.NewOption("No", "n"),
				).
				Value(&addBody),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	if addBody == "y" {
		rl, err := readline.New("> ") // Prompt symbol
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
	}

	reqBody := bytes.NewBuffer([]byte(jsonStr))

	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response\n", string(body))
}

func Delete(url string) {
	resp, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		fmt.Println("Error:", err)
		return

	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print response
	fmt.Println("Response:", string(body))
}
