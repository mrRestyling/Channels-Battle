package main

import "fmt"

func main() {

	one := make(chan int)
	two := make(chan int)

	go func() {
		defer close(one)
		for i := 1; i <= 10; i++ {
			one <- i
		}

	}()

	go func() {
		defer close(two)
		for num := range one {
			two <- num * num
		}

	}()

	for x := range two {
		fmt.Println(x)
	}

}
