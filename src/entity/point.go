package entity

import (
	"log"

	"github.com/lucas625/Projeto-CG/src/utils"
)

// Point is a class for points.
//
// Members:
// 	Coordinates - list of coordinates.
//
type Point struct {
	Coordinates []float64
}

// CheckPointCoordinates is a function to check if two points are of the same size.
//
// Parameters:
// 	p1 - The first vector.
// 	p2 - The second vector.
//
// Returns:
// 	none
//
func CheckPointCoordinates(p1, p2 *Point) {
	if len(p1.Coordinates) != len(p2.Coordinates) {
		log.Fatalf("Invalid size of Point. Expected: %d and Got: %d.\n", len(p1.Coordinates), len(p2.Coordinates))
	}
}

// ExtractVector is a function to extract a vector between two points.
//
// Parameters:
// 	p1 - the starting Point.
//  p2 - the target Point.
//
// Returns:
// 	a Vector.
//
func ExtractVector(p1, p2 *Point) utils.Vector {
	CheckPointCoordinates(p1, p2)
	vaux := utils.InitVector(len(p1.Coordinates))
	for i := 0; i < len(p1.Coordinates); i++ {
		vaux.Coordinates[i] = p2.Coordinates[i] - p1.Coordinates[i]
	}
	return vaux
}

// PointToHomogeneousCoord is a function to add the extra 1 coord and transpose the Point converting it to Matrix.
//
// Parameters:
// 	point - The Point.
//
// Returns:
// 	a Matrix.
//
func PointToHomogeneousCoord(point *Point) utils.Matrix {
	maux := utils.InitMatrix(len(point.Coordinates)+1, 1)
	for i := 0; i < len(point.Coordinates); i++ {
		maux.Values[i][0] = point.Coordinates[i]
	}
	maux.Values[len(point.Coordinates)][0] = 1
	return maux
}

// InitPoint is a function to initialize a Point.
//
// Parameters:
// 	size - The size of the Point.
//
// Returns:
// 	a Point
//
func InitPoint(size int) Point {
	if size < 0 {
		log.Fatalf("Invalid Point size %d.\n", size)
	}
	point := Point{Coordinates: make([]float64, size)}
	return point
}
