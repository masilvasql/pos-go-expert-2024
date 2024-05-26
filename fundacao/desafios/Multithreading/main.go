package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
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

	fastestApiURL := make(chan string)

	go makeCepRequest(ctx, *requestBrasilApi, fastestApiURL)
	go makeCepRequest(ctx, *requestViaCep, fastestApiURL)

	apiUrl := <-fastestApiURL
	cancel()
	fmt.Println(apiUrl)

}

func makeCepRequest(ctx context.Context, requestInfo RequestInfo, fastestApiURL chan string) {
	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, "GET", requestInfo.ApiUrl, nil)
	if err != nil {
		fmt.Printf("Error to create request %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error to make request %v", err)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
	duration := time.Since(start).Milliseconds()
	fastestApiURL <- fmt.Sprintf("The fastest api was %s.\nThe time spen was: %d miliseconds", requestInfo.ApiName, duration)
}
