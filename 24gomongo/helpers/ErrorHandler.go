package helpers

import (
	"fmt"
	"log"
)

func CheckForNilError(e error, m string) {
	if e != nil {
		log.Fatal(e)
	} else {
		fmt.Println(m)
	}
}
