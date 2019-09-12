package utils

type Vector struct {
	/*
		Class for vectors.

		Members:
			X - x coordinate.
			Y - y coordinate.
			Z - z coordinate.

	*/
	X float64
	Y float64
	Z float64
} 

func CMultVector(vect *Vector, k float64) Vector {
	/*
		Function for Scalar Multiplication.

		Parameters:
			vect - The vector.
			k    - The constant for the multiplication.

		Returns:
			The resulting vector.
	*/
	vectAux := Vector{0,0,0}
	vectAux.X = k * vect.X
	vectAux.Y = k * vect.Y
	vectAux.Z = k * vect.Z

	return vectAux
}

func SumVector(vect1 *Vector, vect2 *Vector, k1 float64, k2 float64) Vector {
	/*
		Function to sum 2 vectors.

		Parameters:
			vect1 - The first vector.
			vect2 - The second vector.
			k1    - Constant multiplying the first vector.
			k2    - Constant multiplying the second vector.

		Returns:
			The resulting vector.
	*/
	vect1Aux := CMultVector(vect1, k1)
	vect2Aux := CMultVector(vect2, k2)

	vectAux := Vector{0,0,0}
	vectAux.X = vect1Aux.X + vect2Aux.X
	vectAux.Y = vect1Aux.Y + vect2Aux.Y
	vectAux.Z = vect1Aux.Z + vect2Aux.Z

	return vectAux
}

func DotProduct(vect1 *Vector, vect2 *Vector) float64 {
	/*
		Function to dot product 2 vectors.

		Parameters:
			vect1 - The first vector.
			vect2 - The second vector.

		Returns:
			The resulting sum.
	*/
	return (vect1.X * vect2.X) + (vect1.Y * vect2.Y) + (vect1.Z * vect2.Z)
}

func ProjVector(vect1 *Vector, vect2 *Vector) Vector {
	/*
		Function to project one vector in the other.

		Parameters:
			vect1 - The first vector.
			vect2 - The second vector.

		Returns:
			The resulting vector.
	*/
	topConstant := DotProduct(vect1, vect2)
	bottomConstant := DotProduct(vect2, vect2)

	vectAux := CMultVector(vect2, topConstant / bottomConstant)

	return vectAux

}

func OrtogonalizeVector(vect1 *Vector, vect2 *Vector) Vector {
	/*
		Function to check if two vectors are ortogonal to each other.

		Parameters:
			vect1 - The first vector.
			vect2 - The second vector.

		Returns:
			The resulting vector.
	*/
	vectAux := ProjVector(vect1, vect2)
	return SumVector(vect1, &vectAux, 1, -1)
}

func CheckOrtogonalVector(vect1 *Vector, vect2 *Vector) bool {
	/*
		Function to check if two vectors are ortogonal to each other.

		Parameters:
			vect1 - The first vector.
			vect2 - The second vector.

		Returns:
			A boolean.
	*/
	if DotProduct(vect1, vect2) == 0 {
		return true
	}

	return false
}