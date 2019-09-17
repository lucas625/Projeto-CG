package entity

import (
	"github.com/lucas625/Projeto-CG/src/utils"
)

// Vertices is a class for vertices.
//
// Members:
// 	Vertices - list of vertices.
//
type Vertices struct {
	Points []Point
}

// VerticesToHomogeneousCoord is a function to create a matrix with all vertices in homogeneous coordinates.
//
// Parameters:
// 	vertices - a Vertices.
//
// Returns:
//  the corresponding matrix.
//
func VerticesToHomogeneousCoord(vertices *Vertices) utils.Matrix {
	points := vertices.Points
	maux := utils.InitMatrix(len(points[0].Coordinates)+1, len(points))
	for i, point := range points {
		paux := PointToHomogeneousCoord(&point) // paux is a matrix
		for j := 0; j < len(paux.Values); j++ {
			maux.Values[j][i] = paux.Values[j][0]
		}
	}
	return maux
}

// MatrixToVertices is a function to parse a Matrix into a Vertices (a point is a column, removes the homogeneous coord).
//
// Parameters:
// 	matrix - a Matrix.
//
// Returns:
//  the corresponding Vertices.
//
func MatrixToVertices(matrix *utils.Matrix) Vertices {
	points := make([]Point, len(matrix.Values)-1) //ignoring last value
	for j := 0; j < len(matrix.Values[0]); j++ {
		pointAux := InitPoint(len(matrix.Values) - 1)
		for i := 0; i < len(matrix.Values)-1; i++ {
			pointAux.Coordinates[i] = matrix.Values[i][j]
		}
		points[j] = pointAux
	}
	return InitVertices(points)
}

// MultVertices is a function to multiply all Vertices by a matrix.
//
// Parameters:
// 	vertices - a Vertices.
//  matrix   - the multiplying matrix.
//
// Returns:
// 	the Vertices multiplied by the matrix.
//
func MultVertices(vertices *Vertices, matrix *utils.Matrix) Vertices {
	pointMatrix := VerticesToHomogeneousCoord(vertices)
	maux := utils.MultMatrix(matrix, &pointMatrix)
	vertAux := MatrixToVertices(&maux)
	return vertAux
}

// InitVertices is a function to initialize a Vertices.
//
// Parameters:
// 	points - a list of points.
//
// Returns:
// 	a Vertices.
//
func InitVertices(points []Point) Vertices {
	vertices := Vertices{Points: points}
	return vertices
}
