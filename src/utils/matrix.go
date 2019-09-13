package utils

import (
	"errors"
	"log"
)

// Matrix is a class for matrices.
//
// Members:
// 	Values  - values of the matrix.
// 	Lines   - number of lines of the matrix.
// 	Columns - number of columns of the matrix.
//
type Matrix struct {
	Values  [][]float64
	Lines   int
	Columns int
}

// showError is a function to print an error.
//
// Parameters:
// 	err - The error.
// 	msg - The string with extra informations.
//
// Returns:
// 	none
//
func showError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

// matrixLinesColumns is a function to check the number of lines and columns of a matrix.
//
// Parameters:
// 	lin - The number of lines of the matrix.
// 	col - The number of columns of the matrix.
//
// Returns:
// 	An error.
//
func matrixLinesColumns(lin, col int) error {
	if lin <= 0 || col <= 0 {
		return errors.New("Invalid Number of lines and columns")
	}
	return nil
}

// InitMatrix is a function to initialize a Matrix.
//
// Parameters:
// 	lin - The number of lines of the matrix.
// 	col - The number of columns of the matrix.
//
// Returns:
// 	A Matrix.
//
func InitMatrix(lin, col int) Matrix {
	err := matrixLinesColumns(lin, col)
	showError(err, "lines - %d, columns - %d") //add lin and col here
	matrix := Matrix{Values: make([][]float64, lin, col), Lines: lin, Columns: col}
	return matrix
}
