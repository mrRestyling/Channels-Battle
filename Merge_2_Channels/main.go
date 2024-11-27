package main

import (
	"fmt"
	"time"
)

// Функция func Merge(f func(int) int, in1 <-chan int, in2 <-chan int, out chan <- int, n int)

func main() {

	t := time.Now()

	merge2Channels(func(x int) int { return x * x }, make(chan int), make(chan int), make(chan int), 10)

	fmt.Println(time.Since(t))

}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	type t struct {
		i int
		n int
	}

	ch1 := make(chan t)
	ch2 := make(chan t)

	for i := 0; i < n; i++ {
		go func(num int) {
			x1 := <-in1
			ch1 <- t{i: num, n: fn(x1)}
		}(i)
	}
	for i := 0; i < n; i++ {
		go func(num int) {
			x2 := <-in2
			ch2 <- t{i: num, n: fn(x2)}
		}(i)
	}

	go func() {
		x1sl := make([]t, 0, n)
		x2sl := make([]t, 0, n)

		for i := 0; i < n; i++ {
			x1sl = append(x1sl, <-ch1)
			x2sl = append(x2sl, <-ch2)
		}

		for _, x1 := range x1sl {
			for _, x2 := range x2sl {
				if x1.i == x2.i {
					out <- x1.n + x2.n
					break
				}
			}
		}
	}()
}
