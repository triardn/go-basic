package main

import (
	"fmt"
	"time"
)

func main() {
	go timer("lompat")
	timer("bobo")
}

func timer(todo string) {
	for i := 1; i < 5; i++ {
		fmt.Println(todo, i)
		time.Sleep(time.Second)
	}
}
