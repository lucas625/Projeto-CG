package utils

import (
	"errors"
	"log"
	"strconv"
	"fmt"
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

// PrintMatrix is a function to print a Matrix.
//
// Parameters:
// 	matrix - A point to a Matrix.
//
// Returns:
// 	none
//
func PrintMatrix(matrix *Matrix) {
	totalString := "Matrix - lines: %d columns: %d\n values:\n  "
	for i := range matrix.Values {
		for j := range matrix.Values[i] {
			totalString += strconv.FormatFloat(matrix.Values[i][j], 'g', -1, 64) + " "
		}
		totalString = totalString[:len(totalString)-1] + "\n  "
	}
	fmt.Printf(totalString, matrix.Lines, matrix.Columns)
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
	matrix := Matrix{Values: make([][]float64, lin), Lines: lin, Columns: col}
	// setting default values
	for i := range matrix.Values {
		matrix.Values[i] = make([]float64, col)
	}
	return matrix
}
