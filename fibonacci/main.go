package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonacci(8))

}

func fibonacci(num int) int {

	x := 0

	y := 1

	for i := 0; i < num; i++ {

		x, y = y, x+y

	}

	return y - x

}
