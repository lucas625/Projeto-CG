package main

import (
	"fmt"

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
	//utils.PrintMatrix(&a)
	b := utils.TransposeMatrix(&a)
	//util.PrintMatrix(&b)
	c := utils.MultMatrix(&a, &b)
	//utils.PrintMatrix(&c)
	d := utils.CMultMatrix(&c, 2)
	utils.PrintMatrix(&d)
	v1 := utils.InitVector(3)
	v1.Coordinates = []float64{2, 0, 0}
	v2 := utils.NormalizeVector(&v1)
	fmt.Println(v1, v2)
}
