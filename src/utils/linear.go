package utils

import (
	"log"
)

// InitHomogeneousCoordMatrix is a function to initialize a matrix with homogeneous coordinates.
//
// Parameters:
// 	dim - The number of dimensitons(2 or 3).
//
// Returns:
// 	none
//
func InitHomogeneousCoordMatrix(dim int) Matrix {
	if dim > 3 || dim < 2 {
		log.Fatalf("Invalid dimensions: %d.\n", dim)
	}
	maux := IdMatrix(dim+1)
	return maux
}

// InitTranslationMatrix is a function to initialize a translation matrix.
//
// Parameters:
// 	dim   - The number of dimensitons(2 or 3).
// 	trans - List of values to translate (x,y,z...).
//
// Returns:
// 	none
//
func InitTranslationMatrix(dim int, trans []float64) Matrix {
	if dim > 3 || dim < 2 {
		log.Fatalf("Invalid dimensions: %d.\n", dim)
	} 
	if len(trans) != dim {
		log.Fatalf("Invalid number of values to translate, expected: %d and got: %d.\n", dim, len(trans))
	}
	maux := InitHomogeneousCoordMatrix(dim)
	for i := 0; i < dim; i++ {
		maux.Values[i][dim-1] = trans[i]
	}
	return maux
}