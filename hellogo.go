package main

import (
	"fmt"
)

var pl = fmt.Println

func dblArrVals(arr *[4]int){
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func main() {
	pArr := [4]int{1,2,3,4}
	dblArrVals(&pArr)
	pl(pArr)
}
