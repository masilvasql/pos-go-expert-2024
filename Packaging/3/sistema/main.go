package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/masilvasql/pos-go-expert-2024/Packaging/3/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	fmt.Println(uuid.New().String())
}
