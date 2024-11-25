package main

import "fmt"

// 0 1 1 2 3 5 8 13 21 34

func fib(num int) <-chan int {

	c := make(chan int, num)

	go func(x int, y int) {

		for i := 1; i <= num; i++ {
			x, y = x+y, x
			c <- y
		}
		close(c)

	}(0, 1)

	return c
}

func main() {

	for num := range fib(10) {
		fmt.Println(num)
	}

}
