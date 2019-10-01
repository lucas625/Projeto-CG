package utils

import (
	"errors"
	"fmt"
	"log"
	"strconv"
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
// 	matrix - A pointer to a Matrix.
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

// CMultMatrix is a function to mutiply a Matrix by a constant.
//
// Parameters:
// 	matrix - A pointer to a Matrix.
//  c      - A constant.
//
// Returns:
// 	The resulting matrix
//
func CMultMatrix(matrix *Matrix, c float64) Matrix {
	maux := InitMatrix(matrix.Lines, matrix.Columns)
	for i := 0; i < matrix.Lines; i++ {
		for j := 0; j < matrix.Columns; j++ {
			maux.Values[i][j] += matrix.Values[i][j] * c
		}
	}
	return maux
}

// MultMatrix is a function to mutiply two Matrices.
//
// Parameters:
// 	m1 - A pointer to a Matrix.
//  m2 - A pointer to a Matrix.
//
// Returns:
// 	The resulting matrix
//
func MultMatrix(m1, m2 *Matrix) Matrix {
	if m1.Columns != m2.Lines {
		log.Fatalf("Invalid length of matrices. col m1: %d, lin m2: %d.\n", m1.Columns, m2.Lines)
	}
	maux := InitMatrix(m1.Lines, m2.Columns)
	for i := 0; i < m1.Lines; i++ {
		for j := 0; j < m2.Columns; j++ {
			for z := 0; z < m1.Columns; z++ {
				maux.Values[i][j] += m1.Values[i][z] * m2.Values[z][j]
			}
		}
	}
	return maux
}

// TransposeMatrix is a function to transpose a Matrix.
//
// Parameters:
// 	matrix - the target Matrix.
//
// Returns:
// 	The transposed Matrix.
//
func TransposeMatrix(matrix *Matrix) Matrix {
	maux := InitMatrix(matrix.Columns, matrix.Lines)
	for i := 0; i < matrix.Lines; i++ {
		for j := 0; j < matrix.Columns; j++ {
			maux.Values[j][i] = matrix.Values[i][j]
		}
	}
	return maux
}

// IDMatrix is a function to initialize a Identity Matrix.
//
// Parameters:
// 	size - The number of lines of the matrix.
//
// Returns:
// 	A Matrix.
//
func IDMatrix(size int) Matrix {
	if size < 1 {
		log.Fatalf("Invalid size of identity matrix: %d.\n", size)
	}
	maux := InitMatrix(size, size)
	for i := 0; i < size; i++ {
		maux.Values[i][i] = 1
	}
	return maux
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
	ShowError(err, "lines - %d, columns - %d.") //add lin and col here
	matrix := Matrix{Values: make([][]float64, lin), Lines: lin, Columns: col}
	// setting default values
	for i := range matrix.Values {
		matrix.Values[i] = make([]float64, col)
	}
	return matrix
}
