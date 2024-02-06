package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Price struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	err = SaveInFile(result)
	if err != nil {
		log.Println(err)
	}
}

func SaveInFile(result []byte) error {
	var price Price
	err := json.Unmarshal(result, &price)
	if err != nil {
		return err
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	size, err := fmt.Fprintln(file, "Dólar:", price.Bid)
	if err != nil {
		return err
	}
	fmt.Printf("File created with success! Size: %d bytes\n", size)

	return nil
}
