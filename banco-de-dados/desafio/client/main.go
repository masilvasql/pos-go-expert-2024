package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type DadosCotacao struct {
	Cotacao string `json:"cotacao"`
}

func main() {
	timeOut := 300 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	urlServer := "http://localhost:8080/cotacao"

	chRequest := make(chan DadosCotacao)

	go func() {
		err := makeRequest(ctx, urlServer, chRequest)
		if err != nil {
			fmt.Println("Error to make request ", err)
		}
	}()

	select {
	case <-time.After(timeOut):
		fmt.Println("Error: Request time out limit exceeded")
	case result := <-chRequest:
		err := writeFile(result)
		if err != nil {
			fmt.Println("Error to write in file", err)
		}
	}
}

func writeFile(dados DadosCotacao) error {
	strFormated := fmt.Sprintf("Dólar: %v", dados.Cotacao)

	err := os.WriteFile("cotacao.txt", []byte(strFormated), 0644)
	if err != nil {
		return err
	}

	return nil
}

func makeRequest(ctx context.Context, url string, chRequest chan DadosCotacao) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var cotacao DadosCotacao
	err = json.Unmarshal(resBody, &cotacao)
	if err != nil {
		return err
	}

	chRequest <- cotacao

	return nil
}
