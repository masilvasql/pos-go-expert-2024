package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := range 10 {
		fmt.Printf("%d - %s\n", i, name)
		time.Sleep(time.Second)
		waitGroup.Done()
	}
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(20)

	go task("A", &wg)
	go task("B", &wg)

	wg.Wait()
}
