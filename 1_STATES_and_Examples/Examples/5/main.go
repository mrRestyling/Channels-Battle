package main

import (
	"fmt"
	"sync"
)

// Условие
// Даны n каналов типа int. Нужно написать функцию, которая сольет (смержит) все данные из этих каналов в один и вернет его.

func main() {

	// var wg sync.WaitGroup

	chGo1 := make(chan int)
	chGo2 := make(chan int)
	chGo3 := make(chan int)

	go func() {
		defer close(chGo1)
		chGo1 <- 1
		chGo1 <- 44
	}()
	go func() {
		defer close(chGo2)
		chGo2 <- 2
		chGo2 <- 55
	}()
	go func() {
		defer close(chGo3)
		chGo3 <- 3
		chGo3 <- 66
	}()

	for ch := range merge(chGo1, chGo2, chGo3) {
		fmt.Println(ch)
	}

}

func merge(channels ...<-chan int) <-chan int {

	res := make(chan int)

	var wg sync.WaitGroup

	for _, n := range channels {
		wg.Add(1)

		go func(<-chan int) {
			defer wg.Done()
			for ch := range n {
				res <- ch
			}
		}(n)

	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
