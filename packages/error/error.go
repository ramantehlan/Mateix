package error

import (
	"log"
	"fmt"
)

func Error(err string) {
	fmt.Println("Error: ", err)
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
