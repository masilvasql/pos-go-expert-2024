package main

import (
	"fmt"

	"github.com/masilvasql/massilvautils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
