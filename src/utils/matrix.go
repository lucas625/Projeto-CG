package utils

import "log"

type Matrix struct {
	/*
		Class for matrices.

		Members:
			Values  - values of the matrix.
			Lines   - number of lines of the matrix.
			Columns - number of columns of the matrix.

	*/
	Values []float64
	Lines int
	Columns int
}

func showError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

func matrixLinesColumns(lin, col int) (int, int, error) {
	if lin <= 0 || col <= 0 {
		return lin, col, &argError{"Invalid Number of lines and columns"}
	}
	return 0, 0, nil
}

func InitMatrix(lin, col int) {
	_, _, err := matrixLinesColumns(lin, col)
	showError(err, "lines - %d, columns - %d")//add lin and col here
}