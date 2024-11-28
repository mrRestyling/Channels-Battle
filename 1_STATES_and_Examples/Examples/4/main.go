package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	_, _ = worker(), <-worker()

	fmt.Println(int(time.Since(t).Seconds()))
}

func worker() chan int {

	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	return ch
}
