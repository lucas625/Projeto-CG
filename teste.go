package main

import (
	"fmt"

	"github.com/lucas625/Projeto-CG/src/entity"
)

func main() {
	a := entity.InitPoint(3)
	a.Coordinates[0] = 2
	a.Coordinates[1] = 3
	a.Coordinates[2] = 5
	b := entity.InitPoint(3)
	b.Coordinates[0] = -1
	b.Coordinates[1] = 2
	b.Coordinates[2] = 4
	fmt.Println(entity.PointToHomogeneousCoord(&a), a)
}
