package main

import (
	"fmt"
)

func main() {
	// // basic loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// // while loop
	// x := 0
	// for x < 7 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// // array range loop
	// arr := []int{1, 2, 3}

	// for index, value := range arr {
	// 	fmt.Println("index: ", index, "value: ", value)
	// }

	// map loop
	trashcan := make(map[string]int)

	trashcan["leaf"] = 7
	trashcan["paper"] = 10
	trashcan["plastic"] = 3

	for key, value := range trashcan {
		fmt.Println("key", key, "value", value)
	}
}
