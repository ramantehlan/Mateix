package e

import (
	"fmt"
)

// Error prints the error in the console
func Error(err string) {
	fmt.Println("Error: ", err)
}

// Check check if the error is nil or not, if not then throw the error
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
