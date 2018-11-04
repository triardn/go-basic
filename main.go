package main

import (
	"fmt"
)

func main() {
	stock := make(map[string]int)

	stock["milk"] = 4
	stock["chees"] = 7
	stock["mayonaise"] = 11

	fmt.Println(stock)

	stock["milk"] += 8

	fmt.Println(stock)

	delete(stock, "mayonaise")

	fmt.Println(stock)
}
