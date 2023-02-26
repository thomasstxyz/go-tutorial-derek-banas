package main

import (
	"fmt"
)

var pl = fmt.Println

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	vArr := []int{1,2,3,4}
	pl("Array Sum :", getArraySum(vArr))

	
}
