package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type RequestInfo struct {
	ApiName string
	ApiUrl  string
}

func NewRequestInfo(apiName, apiUrl string) *RequestInfo {
	return &RequestInfo{
		ApiName: apiName,
		ApiUrl:  apiUrl,
	}
}

func main() {

	cep := "89201215"

	requestBrasilApi := NewRequestInfo("Brasil API", fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))

	requestViaCep := NewRequestInfo("Via CEP", fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	responseChanRequest := make(chan string)

	go makeCepRequest(ctx, *requestBrasilApi, responseChanRequest)
	go makeCepRequest(ctx, *requestViaCep, responseChanRequest)

	select {
	case <-time.After(time.Second):
		log.Println("timeout exceeded")
		cancel()
	case result := <-responseChanRequest:
		fmt.Println(result)
		cancel()
	}

}

func makeCepRequest(ctx context.Context, requestInfo RequestInfo, responseChanRequest chan string) {
	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, "GET", requestInfo.ApiUrl, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	jsonResult, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	duration := time.Since(start).Milliseconds()
	responseChanRequest <- fmt.Sprintf("The fastest api was %s.\nThe spent time was: %d miliseconds. the request body is\n%s", requestInfo.ApiName, duration, jsonResult)
}
