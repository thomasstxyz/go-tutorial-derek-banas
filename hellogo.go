package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	intSum := func(x, y int) int { return x + y}
	pl("5 + 4 =", intSum(5, 4))

	samp1 := 1
	changeVar := func(){
		samp1 += 1
	}
	changeVar()
	pl("samp1 =", samp1)
}
