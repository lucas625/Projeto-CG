package main

import (
	"fmt"

	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/utils"
)

func main() {
	a := entity.InitPoint(3)
	a.Coordinates[0] = 0
	a.Coordinates[1] = 0
	a.Coordinates[2] = 1
	b := entity.InitPoint(3)
	b.Coordinates[0] = 0
	b.Coordinates[1] = 0
	b.Coordinates[2] = 2
	c := entity.InitPoint(3)
	c.Coordinates[0] = 5
	c.Coordinates[1] = 3
	c.Coordinates[2] = 2
	// matrix
	maux := utils.InitTranslationMatrix(3, []float64{10, 15, 25})
	maux.Values[0][1] = 2
	maux.Values[0][2] = -3
	maux.Values[1][0] = -1
	maux.Values[1][2] = 5
	maux.Values[2][0] = 1
	maux.Values[2][1] = 1
	// vertices
	vertices := entity.InitVertices([]entity.Point{a, b, c})
	vaux := entity.MultVertices(&vertices, &maux)
	utils.PrintMatrix(&maux)
	fmt.Println(vaux)
	// triangle
	triang := entity.InitTriangle([]int{0, 1, 2})
	fmt.Println(triang)
	cam := camera.InitCameraWithPoints(&a, &b)
	fmt.Println(cam)
}
