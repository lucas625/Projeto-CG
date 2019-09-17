package main

import (
	"github.com/lucas625/Projeto-CG/src/utils"
)

func main() {
	a := utils.InitTranslationMatrix(3, []float64{-2, 10, -3})
	a.Values[0][1] = 2
	a.Values[0][2] = 20
	a.Values[1][0] = -4
	a.Values[1][2] = 30
	a.Values[2][0] = -2
	a.Values[2][1] = 45
	utils.PrintMatrix(&a)
	b := utils.TransposeMatrix(&a)
	utils.PrintMatrix(&b)
	c := utils.MultMatrix(&a, &b)
	utils.PrintMatrix(&c)
	d := utils.CMultMatrix(&c, 2)
	utils.PrintMatrix(&d)
}
