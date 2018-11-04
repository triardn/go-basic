package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// sum function
	result := sum(17.4, 11.8)
	fmt.Println(result)

	// try error handling
	sqrtResult, errors := sqrt(-90)

	if errors != nil {
		fmt.Println(errors)
	} else {
		fmt.Println(sqrtResult)
	}
}

func sum(x float64, y float64) float64 {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("harus lebih besar dari 0")
	}

	return math.Sqrt(x), nil
}
