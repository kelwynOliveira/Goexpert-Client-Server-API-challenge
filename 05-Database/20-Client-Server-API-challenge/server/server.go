package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type QuoteAPI struct {
	USDBRL Quote `json:"USDBRL"`
}

type Quote struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Price struct {
	Bid string `json:"bid"`
}

func main() {
	http.HandleFunc("/cotacao", HandleQuote)
	http.ListenAndServe(":8080", nil)
}

func HandleQuote(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Look into API
	quotation, err := SearchUSDBRL()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Saves on DataBase

	err = DataBase(quotation)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Returns the Bid in json format
	var bid Price
	bid.Bid = quotation.USDBRL.Bid
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bid)
}

func SearchUSDBRL() (*QuoteAPI, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var quote QuoteAPI
	err = json.Unmarshal(result, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

func NewQuote(quote *QuoteAPI) *Quote {
	return &Quote{
		Code:       quote.USDBRL.Code,
		Codein:     quote.USDBRL.Codein,
		Name:       quote.USDBRL.Name,
		High:       quote.USDBRL.High,
		Low:        quote.USDBRL.Low,
		VarBid:     quote.USDBRL.VarBid,
		PctChange:  quote.USDBRL.PctChange,
		Bid:        quote.USDBRL.Bid,
		Ask:        quote.USDBRL.Ask,
		Timestamp:  quote.USDBRL.Timestamp,
		CreateDate: quote.USDBRL.CreateDate,
	}
}

func DataBase(quote *QuoteAPI) error {
	db, err := gorm.Open(sqlite.Open("./quotes.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Quote{})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	usdbrl := NewQuote(quote)

	return db.WithContext(ctx).Create(&usdbrl).Error
}
