package utils

import (
	"log"
)

// Vector is a class for vectors.
//
// Members:
// 	Coordinates - list of coordinates.
//
type Vector struct {
	Coordinates []float64
}

// CheckVectorCoordinates is a function to check if two vectors are of the same size.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	none
//
func CheckVectorCoordinates(vect1 *Vector, vect2 *Vector) {
	if len(vect1.Coordinates) != len(vect2.Coordinates) {
		log.Fatalf("Invalid size of vector. Expected: %d and Got: %d.\n", len(vect1.Coordinates), len(vect2.Coordinates))
	}
}

// CMultVector is a function for Scalar Multiplication.
//
// Parameters:
// 	vect - The vector.
// 	k    - The constant for the multiplication.
//
// Returns:
// 	The resulting vector.
//
func CMultVector(vect *Vector, k float64) Vector {
	vectAux := InitVector(len(vect.Coordinates))
	for i := 0; i < len(vect.Coordinates); i++ {
		vectAux.Coordinates[i] = k * vect.Coordinates[i]
	}
	return vectAux
}

// SumVector is a function to sum 2 vectors.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
// 	k1    - Constant multiplying the first vector.
// 	k2    - Constant multiplying the second vector.
//
// Returns:
// 	The resulting vector.
//
func SumVector(vect1 *Vector, vect2 *Vector, k1 float64, k2 float64) Vector {
	CheckVectorCoordinates(vect1, vect2)
	vect1Aux := CMultVector(vect1, k1)
	vect2Aux := CMultVector(vect2, k2)

	vectAux := InitVector(len(vect1.Coordinates))
	for i := 0; i < len(vect1.Coordinates); i++ {
		vectAux.Coordinates[i] = vect1Aux.Coordinates[i] + vect2Aux.Coordinates[i]
	}
	return vectAux
}

// DotProduct is a function to dot product 2 vectors.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting sum.
//
func DotProduct(vect1 *Vector, vect2 *Vector) float64 {
	CheckVectorCoordinates(vect1, vect2)

	var totalSum float64
	for i := 0; i < len(vect1.Coordinates); i++ {
		totalSum += vect1.Coordinates[i] * vect2.Coordinates[i]
	}
	return totalSum
}

// ProjVector is a function to project one vector in the other.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting vector.
//
func ProjVector(vect1 *Vector, vect2 *Vector) Vector {
	CheckVectorCoordinates(vect1, vect2)
	topConstant := DotProduct(vect1, vect2)
	bottomConstant := DotProduct(vect2, vect2)

	vectAux := CMultVector(vect2, topConstant/bottomConstant)

	return vectAux

}

// OrtogonalizeVector is a function to ortogonalize two vectors.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting vector.
//
func OrtogonalizeVector(vect1 *Vector, vect2 *Vector) Vector {
	vectAux := ProjVector(vect1, vect2)
	return SumVector(vect1, &vectAux, 1, -1)
}

// CheckOrtogonalVector is a function to check if two vectors are ortogonal to each other.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	A boolean.
//
func CheckOrtogonalVector(vect1 *Vector, vect2 *Vector) bool {
	if DotProduct(vect1, vect2) == 0 {
		return true
	}

	return false
}

// VectorToHomogeneousCoord is a function to add the extra 0 coord and transpose the Vector converting it to Matrix.
//
// Parameters:
// 	vect - The Vector.
//
// Returns:
// 	a Matrix.
//
func VectorToHomogeneousCoord(vect *Vector) Matrix {
	maux := InitMatrix(len(vect.Coordinates)+1, 1)
	for i := 0; i < len(vect.Coordinates); i++ {
		maux.Values[i][0] = vect.Coordinates[i]
	}
	return maux
}

// VectorCrossProduct is a function to calculate the cross product of twp Vectors.
//
// Parameters:
// 	vect1 - The first Vector.
//  vect2 - The second Vector.
//
// Returns:
// 	a Vector
//
func VectorCrossProduct(vect1, vect2 *Vector) Vector {
	CheckVectorCoordinates(vect1, vect2)
	if len(vect1.Coordinates) != 3 {
		log.Fatalf("Invalid size of 3D vector: %d.\n", len(vect1.Coordinates))
	}
	vaux := InitVector(3)
	coord1 := vect1.Coordinates
	coord2 := vect2.Coordinates
	i := (coord1[1] * coord2[2]) -(coord1[2] * coord2[1])
	j := (coord1[2] * coord2[0]) -(coord1[0] * coord2[2])
	k := (coord1[0] * coord2[1]) -(coord1[1] * coord2[0])
	vaux.Coordinates = []float64{i, j, k}
	return vaux
}

// InitVector is a function to initialize a Vector.
//
// Parameters:
// 	size - The size of the Vector.
//
// Returns:
// 	a Vector
//
func InitVector(size int) Vector {
	if size < 0 {
		log.Fatalf("Invalid vector size %d.\n", size)
	}
	vect := Vector{Coordinates: make([]float64, size)}
	return vect
}