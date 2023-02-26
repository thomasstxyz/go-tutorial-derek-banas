package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func changeVal2(myPtr *int) int {
	*myPtr = 12
	return *myPtr
}

func main() {
	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)
	pl()

	f4 := 5
	pl("f4 before func :", f4)
	changeVal2(&f4)
	pl("f4 after func :", f4)
	pl()
}
