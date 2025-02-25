package utils

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) {
	resp, err := http.Get(url)

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
