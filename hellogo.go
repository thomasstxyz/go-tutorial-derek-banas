package main

import (
	"fmt"
)

var pl = fmt.Println

func getSum(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}

func main() {
	// func funcName(parameters) returnType {BODY}

	pl(getSum(5, 4))
	pl(getTwo(2))
	pl(getQuotient(5, 0))
	pl(getQuotient(5, 4))
}
