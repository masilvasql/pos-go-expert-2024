package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Cotacao struct {
	Usdbrl struct {
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
	} `json:"USDBRL"`
}

type DadosCotacao struct {
	Id  int `gorm:"primaryKey"`
	Bid string
	gorm.Model
}

type DadosCotacaoOutput struct {
	Cotacao string `json:"cotacao"`
}

var db *gorm.DB

func main() {

	http.HandleFunc("/cotacao", handleGetCotacao)
	fmt.Println("SERVER IS RUNING")
	setupDb()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func setupDb() {
	var err error
	db, err = gorm.Open(sqlite.Open("cotacao_db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is confifured")
	db.AutoMigrate(&DadosCotacao{})
}

func handleGetCotacao(w http.ResponseWriter, r *http.Request) {
	timeout := 200 * time.Millisecond

	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	urlEconomia := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	chRequest := make(chan Cotacao)

	go func() {
		err := makeRequest(ctx, urlEconomia, chRequest)
		if err != nil {
			fmt.Println("Error making request:", err)
		}
	}()

	select {
	case <-time.After(timeout):
		fmt.Println("Error: Request time out")
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Request Time out"))
		cancel()
	case result := <-chRequest:
		insertDataOnDataBase(&result)

		output := &DadosCotacaoOutput{
			Cotacao: result.Usdbrl.Bid,
		}

		data, err := json.Marshal(output)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		ctx.Done()
	}

}

func makeRequest(ctx context.Context, url string, chCotacao chan Cotacao) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var cotacao Cotacao

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, &cotacao)

	if err != nil {
		return err
	}

	chCotacao <- cotacao

	return nil
}

func insertDataOnDataBase(cotacao *Cotacao) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	done := make(chan bool)

	go func() {
		dadosCotacao := &DadosCotacao{
			Bid: cotacao.Usdbrl.Bid,
		}

		if err := db.WithContext(ctx).Create(&dadosCotacao).Error; err != nil {
			log.Printf("Error inserting data: %v", err)
			done <- false
		} else {
			done <- true
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("ERROR: TIMEOUT EXCEEDED TO INSERT DATA INTO DATABASE")
	case success := <-done:
		if success {
			fmt.Println("Dados inseridos antes do timeout")
		} else {
			fmt.Println("ERROR: Failed to insert data")
		}
	}
}
