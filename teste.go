package main

import (
	"fmt"
	"github.com/lucas625/Projeto-CG.git/src/utils"
)

func main() {
	a := vector.Vector{1,2,3}
	b := vector.CMultVector(&a, 3)

	fmt.Println(a, b)
}