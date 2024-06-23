package main

func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}

func recebe(nome string, hello chan<- string) {
	hello <- nome
}

func ler(hello <-chan string) {
	println(<-hello)
}
