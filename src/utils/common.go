package utils

import (
	"log"
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
