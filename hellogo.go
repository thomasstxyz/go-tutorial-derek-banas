package main

import (
	"fmt"
)

var pl = fmt.Println
func main() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)
}
