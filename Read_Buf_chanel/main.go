package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	// c <- 4
	// iteration terminates after receiving 3 values
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}

// Пул воркеров https://habr.com/ru/articles/490336/
