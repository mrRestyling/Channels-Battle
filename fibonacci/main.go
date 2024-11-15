package main

import "fmt"

func main() {
	ch := make(chan int)
	q := make(chan struct{})

	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}

		q <- struct{}{}

	}()

	fibonacci(ch, q)
}

func fibonacci(ch chan<- int, quit <-chan struct{}) {

	x, y := 0, 1

	defer close(ch)

	for {
		select {
		case ch <- x:
			x, y = y, x+y

		case <-quit:
			return
		}

	}

}
