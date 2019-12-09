package utils

import (
	"log"
	"os"
)

// ShowError is a function to print an error.
//
// Parameters:
// 	err - The error.
// 	msg - The string with extra informations.
//
// Returns:
// 	none
//
func ShowError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

// PathExists returns a boolean checking if the target file or folder exists.
//
// Parameters:
// 	path - Path to the file or directory.
//
// Returns:
// 	a boolean.
//
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CheckTolerance is a function to check the tolerance.
//
// Parameters:
// 	value - the target value.
//  pos   - the value it must be compared with.
//
// Returns:
// 	a flag if within the tolerance.
//
func CheckTolerance(value, pos float64) bool {
	tol := 10E-10
	return (pos-tol <= value && value <= pos+tol)
}
