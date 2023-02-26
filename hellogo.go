package main

import (
	"fmt"
)

var pl = fmt.Println
func main() {
	// for loop
	// for initialization; condition; postStatement {BODY}
	for x := 1; x <= 5; x++ {
		pl(x)
	}

	// while loop
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}
}
