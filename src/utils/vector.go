package utils

import (
	"log"
)

// Vector is a class for vectors.
//
// Members:
// 	X - x coordinate.
// 	Y - y coordinate.
// 	Z - z coordinate.
//
type Vector struct {
	X float64
	Y float64
	Z float64
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
	vectAux := Vector{0, 0, 0}
	vectAux.X = k * vect.X
	vectAux.Y = k * vect.Y
	vectAux.Z = k * vect.Z

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
	vect1Aux := CMultVector(vect1, k1)
	vect2Aux := CMultVector(vect2, k2)

	vectAux := Vector{0, 0, 0}
	vectAux.X = vect1Aux.X + vect2Aux.X
	vectAux.Y = vect1Aux.Y + vect2Aux.Y
	vectAux.Z = vect1Aux.Z + vect2Aux.Z

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
	return (vect1.X * vect2.X) + (vect1.Y * vect2.Y) + (vect1.Z * vect2.Z)
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

// VectorToList is a function to convert the vector to a list.
//
// Parameters:
// 	vect - The vector.
//
// Returns:
// 	The converted vector as list.
//
func VectorToList(vect *Vector) []float64 {
	return []float64{vect.X, vect.Y, vect.Z}
}

func InitVector(size int) {
	if size < 0 {
		log.Fatalf("Invalid vector size %d.\n", size)
	}
}