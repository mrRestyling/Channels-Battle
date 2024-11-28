package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// run()
	run2()
}

func run() {

	ch := make(chan int)

	for i := 0; i < 3; i++ {

		go func(idx int) {

			ch <- (idx + 1) * 2

		}(i)

	}

	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}

func run2() {

	wg := &sync.WaitGroup{}

	ch := make(chan int, 3)

	for i := 0; i < 3; i++ {

		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			ch <- (idx + 1) * 2

		}(i)

	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for a := range ch {
		fmt.Println(a)
	}

	// fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}
