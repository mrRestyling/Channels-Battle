package main

import (
	"fmt"
	"time"
)

func main() {

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}

	fmt.Println("[main] Wrote 5 tasks")

	close(tasks)

	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("STOP")

}

func sqrWorker(tasks <-chan int, results chan<- int, id int) {

	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] Sending result by worker %v\n", id, id)
		results <- num * num * num
	}
}
