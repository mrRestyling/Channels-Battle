package main

import (
	"fmt"
	"sync"
)

// Условие
// Даны n каналов типа int. Нужно написать функцию, которая сольет (смержит) все данные из этих каналов в один и вернет его.

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	resch := make(chan int)

	go func() {
		ch1 <- 11
		ch1 <- 555
		close(ch1)
	}()

	go func() {
		ch2 <- 22
		ch2 <- 777
		close(ch2)
	}()

	go func() {
		ch3 <- 33
		ch3 <- 999
		close(ch3)
	}()

	go merge(resch, ch1, ch2, ch3)

	for ch := range resch {
		fmt.Println(ch)
	}

}

func merge(result chan<- int, channels ...chan int) {

	var wg sync.WaitGroup

	for _, c := range channels {

		wg.Add(1)
		go func(chan int) {
			defer wg.Done()
			for num := range c {
				result <- num
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
}

// func main() {

// 	chGo1 := make(chan int)
// 	chGo2 := make(chan int)
// 	chGo3 := make(chan int)

// 	go func() {
// 		defer close(chGo1)
// 		chGo1 <- 1
// 		chGo1 <- 44
// 	}()
// 	go func() {
// 		defer close(chGo2)
// 		chGo2 <- 2
// 		chGo2 <- 55
// 	}()
// 	go func() {
// 		defer close(chGo3)
// 		chGo3 <- 3
// 		chGo3 <- 66
// 	}()

// 	for ch := range merge(chGo1, chGo2, chGo3) {
// 		fmt.Println(ch)
// 	}

// }

// func merge(channels ...<-chan int) <-chan int {

// 	res := make(chan int)

// 	var wg sync.WaitGroup

// 	for _, n := range channels {
// 		wg.Add(1)

// 		go func(<-chan int) {
// 			defer wg.Done()
// 			for ch := range n {
// 				res <- ch
// 			}
// 		}(n)

// 	}

// 	go func() {
// 		wg.Wait()
// 		close(res)
// 	}()

// 	return res
// }
