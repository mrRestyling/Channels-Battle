package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sqr(&wg, tasks, results, i)
	}

	for i := 1; i <= 5; i++ {
		tasks <- i * 2
	}

	fmt.Println("go 5 tasks")

	close(tasks)

	wg.Wait()

	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Println(i, " + ", result)
	}

	fmt.Println("STOP")
}

func sqr(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, insta int) {

	defer wg.Done()
	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Println("--- --- ---", insta)
		results <- num * num

	}

}
