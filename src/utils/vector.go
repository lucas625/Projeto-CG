package utils

type Vector struct {
	/*
		Class for vectors.

		Members:
			x - x coordinate.
			y - y coordinate.
			z - z coordinate.

	*/
	x float64
	y float64
	z float64
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
	vectAux.x = k * vect.x
	vectAux.y = k * vect.y
	vectAux.z = k * vect.z

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
	vectAux.x = vect1Aux.x + vect2Aux.x
	vectAux.y = vect1Aux.y + vect2Aux.y
	vectAux.z = vect1Aux.z + vect2Aux.z

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
	return (vect1.x * vect2.x) + (vect1.y * vect2.y) + (vect1.z * vect2.z)
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