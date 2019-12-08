package main

import (
	"fmt"

	"github.com/lucas625/Projeto-CG/src/camera"
	"github.com/lucas625/Projeto-CG/src/entity"
	"github.com/lucas625/Projeto-CG/src/utils"
)

func main() {

	var p1, p2, p3, p4 entity.Point

	p1.Coordinates = []float64{1, 1, 1}
	p2.Coordinates = []float64{1, 2, 3}
	p3.Coordinates = []float64{-1, -3, -5}
	p4.Coordinates = []float64{0, -2, 3}

	pointL := []entity.Point{p1, p2, p3, p4}

	cameraPath := "resources/json/camera.json"
	cam := camera.LoadJSONCamera(cameraPath)
	fmt.Println("cam:\n", cam)
	camMatrix := camera.CamToHomogeneousMatrix(cam)
	fmt.Println("cam M:\n", camMatrix)

	vertices := entity.Vertices{Points: pointL}
	fmt.Println("vertices:\n", vertices)

	verticesMatrix := entity.VerticesToHomogeneousCoord(&vertices)
	fmt.Println("matrix vertices:\n", verticesMatrix)

	resultMatrix := utils.MultMatrix(&camMatrix, &verticesMatrix)
	fmt.Println("resultMatrix:\n", resultMatrix)

	vaux := entity.MatrixToVertices(&resultMatrix)
	fmt.Println("pos vertices:\n", vaux)

	//
	// obj.Vertices = entity.MatrixToVertices(&resultMatrix)
}
