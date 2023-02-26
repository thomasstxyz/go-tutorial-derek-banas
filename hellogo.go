package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal2(myPtr *int) int {
	*myPtr = 12
	return *myPtr
}

func main() {
	f4 := 10
	var f4Ptr *int = &f4
	pl("f4 Address :", f4Ptr)
	pl("f4 Value :", *f4Ptr)
	*f4Ptr = 11
	pl("f4 Value :", *f4Ptr)

	pl("f4 before func :", f4)
	changeVal2(&f4)
	pl("f4 after func :", f4)
}
