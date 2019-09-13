package main

import (
	"fmt"

	"github.com/lucas625/Projeto-CG/src/utils"
)

func main() {
	a := utils.InitMatrix(3, 4)
	a.Values[1][2] = 10
	fmt.Println(a)
	utils.PrintMatrix(&a)
}
