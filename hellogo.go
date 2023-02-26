package main

import (
	"fmt"
)

var pl = fmt.Println

type Tsp float64
type TBs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBsToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {
	// ml1 := ML(Tsp(3) * 4.92)
	// fmt.Printf("3 tsps = %.2f ML\n", ml1)
	// ml2 := ML(TBs(3) * 14.79)
	// fmt.Printf("3 TBs = %.2f ML\n", ml2)
	// pl("2 tsp + 4 tsp =", Tsp(2) + Tsp(4))
	// pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))
	// fmt.Printf("3 tsp = %.2f ML\n", tspToML(3))
	// fmt.Printf("3 tsp = %.2f ML\n", TBsToML(3))

	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))

	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())

}
