package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println
func main() {
	//            int, float64,  bool, string, rune
	// Default      0,     0.0, false,     "",   ""

	pl(reflect.TypeOf(15))
	pl(reflect.TypeOf(13.3))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("hello"))
	pl(reflect.TypeOf('🦧'))
}
