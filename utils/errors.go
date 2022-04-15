package utils

import "fmt"

// LogErrors takes variadic arguments for errors and format-prints them
func LogErrors(errors ...error) {
	if len(errors) == 0 {
		return
	}

	for _, err := range errors {
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
