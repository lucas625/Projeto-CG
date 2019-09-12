package main

import (
	"fmt"
	"github.com/lucas625/Projeto-CG/src/utils"
)

func main() {
	a := utils.Vector{1,2,3}
	b := utils.CMultVector(&a, 3)

	fmt.Println(a, b)
}