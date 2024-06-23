package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	qtdWorkers := 1000

	for w := range qtdWorkers {
		go worker(w, data)
	}

	for i := range 10000 {
		data <- i
	}
}

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}
