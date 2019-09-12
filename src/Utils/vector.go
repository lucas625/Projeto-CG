package main

import (
	"fmt"
)

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

func SumVector(vect1 *Vector, vect2 *Vector) Vector {
	/*
		Function to sum 2 vectors.

		Parameters:
			vect1 - The first vector.
			vect2 - The second vector.

		Returns:
			The resulting vector.
	*/

	vectAux := Vector{0,0,0}
	vectAux.x = vect1.x + vect2.x
	vectAux.y = vect1.y + vect2.y
	vectAux.z = vect1.z + vect2.z

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

func main() {
	a := Vector{10,20,30}
	b := Vector{2,4,8}
	c := CMultVector(&b, 3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	d := SumVector(&a, &b)

	fmt.Println(a, b, c, d)
}