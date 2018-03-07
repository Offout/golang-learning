package main

import (
	"fmt"
	"math"
)


func Sqrt(number float64) float64 {
	if number < 0 {
		return math.NaN()
	}
	if number == 0 {
		return 0.0
	}
	x1 := number/2
	var x2 float64
	for  {
		x2 = x1
		x1 = (x2 + (number / x2)) / 2
		if (x1 - x2) == 0 {
			break
		}
	}
	return x2
}

func main() {
	fmt.Println(Sqrt(0))
	fmt.Println(math.Sqrt(0))
}
