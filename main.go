package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Tri", Age: 26}

	fmt.Println(p.Name)

	// Array of struct
	arrp := []Person{
		{Name: "Tri", Age: 26},
		{Name: "Tristan", Age: 3},
	}

	fmt.Println(arrp[1].Name)

	// Traverse over array of struct
	for _, elem := range arrp {
		fmt.Println(elem.Name)
		fmt.Println(elem.Age)
		fmt.Println("")
	}
}
