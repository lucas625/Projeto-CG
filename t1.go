package main

import (
	"fmt"

	"github.com/lucas625/Projeto-CG/src/entity"
)

func main() {
	p1 := entity.InitPoint(3)
	p2 := entity.InitPoint(3)
	p3 := entity.InitPoint(3)
	p1.Coordinates = []float64{1, 0, 0}
	p2.Coordinates = []float64{0, 1, 0}
	p3.Coordinates = []float64{0, 0, 1}
	ptarget := entity.InitPoint(3)
	ptarget.Coordinates = []float64{2, 0, 0}
	coordinates := entity.FindBaricentricCoordinates([]entity.Point{p1, p2, p3}, ptarget)
	fmt.Println(coordinates)
}
