package entity

import (
	"github.com/lucas625/Projeto-CG/src/utils"
)

// Vertices is a class for vertices.
//
// Members:
// 	VertList - list of all vertices.
//
type Vertices struct {
	VertList []Vertex
}

// Vertex is a class for a single vertex.
//
// Members:
// 	Point  - position of the vertex.
//  Normal - index of the normal vector.
//
type Vertex struct {
	Point  Point
	Normal int
}

// VerticesToHomogeneousCoord is a function to create a matrix with all vertices in homogeneous coordinates.
//
// Parameters:
// 	vertices - a Vertices.
//
// Returns:
//  the corresponding matrix.
//
func VerticesToHomogeneousCoord(vertices *Vertices) (*utils.Matrix, *[]int) {
	leng := len(vertices.VertList)
	points := make([]Point, leng)
	normals := make([]int, leng)
	for i, vert := range vertices.VertList {
		points[i] = vert.Point
		normals[i] = vert.Normal
	}
	maux := utils.InitMatrix(len(points[0].Coordinates)+1, len(points))
	for i, point := range points {
		paux := PointToHomogeneousCoord(&point) // paux is a matrix
		for j := 0; j < len(paux.Values); j++ {
			maux.Values[j][i] = paux.Values[j][0]
		}
	}
	return &maux, &normals
}

// MatrixToVertices is a function to parse a Matrix into a Vertices (a point is a column, removes the homogeneous coord).
//
// Parameters:
// 	matrix     - a Matrix.
//  normalList - list with normal vectors indices.
//
// Returns:
//  the corresponding Vertices.
//
func MatrixToVertices(matrix *utils.Matrix, normalList *[]int) Vertices {
	vertList := make([]Vertex, len(matrix.Values)-1) //ignoring last value
	for j := 0; j < len(matrix.Values[0]); j++ {
		pointAux := InitPoint(len(matrix.Values) - 1)
		for i := 0; i < len(matrix.Values)-1; i++ {
			pointAux.Coordinates[i] = matrix.Values[i][j]
		}
		vertList[j] = Vertex{Point: pointAux, Normal: (*normalList)[j]}
	}
	return InitVertices(vertList)
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
	pointMatrix, normals := VerticesToHomogeneousCoord(vertices)
	maux := utils.MultMatrix(matrix, pointMatrix)
	vertAux := MatrixToVertices(&maux, normals)
	return vertAux
}

// InitVertices is a function to initialize a Vertices.
//
// Parameters:
// 	vertList - a list of vertices.
//
// Returns:
// 	a Vertices.
//
func InitVertices(vertList []Vertex) Vertices {
	vertices := Vertices{VertList: vertList}
	return vertices
}
