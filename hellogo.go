package main

import (
	"fmt"
)

var pl = fmt.Println

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	pl(getSum2(1,2,3,4,44,3,2,25,2))
}
