package client

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func Client() {
	// Using the "context" package, client.go will have a maximum timeout of 300ms to receive the result from server.go.
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// The client.go should make an HTTP request to server.go asking for the dollar exchange rate.
	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// The client.go will only need to receive the current exchange rate from the server.go (JSON "bid" field).
	io.Copy(os.Stdout, response.Body)

	// The client.go will have to save the current exchange rate in a "cotacao.txt" file in the following format: Dollar: {value}
	// file, err := os.Create("cotacao.txt")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error creating the file: %v.\n", err)
	// }
	// defer file.Close()
	// _, err := file.WriteString(fmt.Sprintf("Dollar: %v", response.Body))

}
