package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		atomic.AddUint64(&number, 1)

		w.Write([]byte(fmt.Sprintf("O total de visitas é de %d", number)))
	})

	http.ListenAndServe(":3003", nil)
}
