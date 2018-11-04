package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	arr[1] = 4

	// short hand fix length array
	arr1 := [3]int{2, 3, 4}

	// short hand dynamic length array
	arr2 := []int{4, 5, 6}
	arr2 = append(arr2, 7)

	fmt.Println(arr1)
	fmt.Println(arr2)
}
