package main

import (
	"fmt"
)

func main() {
	a := vector.Vector{1,2,3}
	b := CMultVector(&a, 3)

	fmt.Println(a, b)
}