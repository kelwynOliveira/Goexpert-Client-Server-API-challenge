package client

import (
	"fmt"
	"net/http"
	"os"
)

// The client.go should make an HTTP request to server.go asking for the dollar exchange rate.
func Request() {
	request, err := http.Get("http://localhost:8080/cotacao")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error accessing %v.\n", err)
	}
	defer request.Body.Close()
}

// The client.go will have to save the current exchange rate in a "cotacao.txt" file in the following format: Dollar: {value}
func CreatingFile() {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating the file: %v.\n", err)
	}
	file.Close()
}

func SavingInFile(fileToWrite string) {
	file, err := os.Open(fileToWrite)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Opening the file: %v.\n", err)
	}
	defer file.Close()

	_, err := file.WriteString(fmt.Sprintf("Dollar: %v", data.bide))
}
