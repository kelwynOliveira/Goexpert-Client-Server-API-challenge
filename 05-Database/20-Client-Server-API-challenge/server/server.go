package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// type Quotation struct {
// 	Usdbrl struct {
// 		Code       string `json:"code"`
// 		Codein     string `json:"codein"`
// 		Name       string `json:"name"`
// 		High       string `json:"high"`
// 		Low        string `json:"low"`
// 		VarBid     string `json:"varBid"`
// 		PctChange  string `json:"pctChange"`
// 		Bid        string `json:"bid"`
// 		Ask        string `json:"ask"`
// 		Timestamp  string `json:"timestamp"`
// 		CreateDate string `json:"create_date"`
// 	} `json:"USDBRL"`
// }

func Server() {
	http.HandleFunc("/cotacao", SearchBid)
	http.ListenAndServe(":8080", nil)
}

func SearchBid(w http.ResponseWriter, r *http.Request) {
	request := RequestJSON()
	result := ReadJSON(request)
	StructureCreation(result, request)

	w.Write([]byte("Hello, World!"))
}

// server.go must consume the API containing the Dollar/Real exchange rate at the address: https://economia.awesomeapi.com.br/json/last/USD-BRL and then return the result to the client in JSON format.

func RequestJSON() *http.Response {
	request, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error accessing %v.\n", err)
	}
	// fmt.Printf("request Type %T \n", request)
	return request
}

func ReadJSON(request *http.Response) []uint8 {
	result, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading result: %v.\n", err)
	}
	return result
}

func StructureCreation(result []uint8, request *http.Response) {
	var data Quotation
	err := json.Unmarshal(result, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing result: %v.\n", err)
	}
	fmt.Println(data)

	request.Body.Close()
}
