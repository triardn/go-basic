package main

import (
	"fmt"
)

func main() {
	a := 7
	if a < 4 {
		fmt.Println("Kurang dari 4")
	} else if a == 4 {
		fmt.Println("Sama dengan 4")
	} else {
		fmt.Println("Lebih dari 4")
	}
}
