package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println
func main() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)
}
