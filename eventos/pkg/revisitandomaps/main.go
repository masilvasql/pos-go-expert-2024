package main

import "fmt"

func main() {
	evento := []string{"teste", "teste2", "teste4", "teste4"}
	fmt.Println(evento[:0])
	fmt.Println(evento[:2])
	evento = append(evento[:0], evento[1:]...)
	fmt.Println(evento)

	evento = append(evento, "Outro Valor")
	fmt.Println(evento)
}
