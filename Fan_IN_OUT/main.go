package main

import (
	"fmt"
	"sync"
)

func main() {

	chanInputNums := getInputChan()

	chanOptSqr1 := getSquareChan(chanInputNums)
	chanOptSqr2 := getSquareChan(chanInputNums)

	chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)

	sqrSum := 0

	for num := range chanMergedSqr {
		sqrSum += num
	}

	fmt.Println("Sum of squares between 0-9 is", sqrSum)
}

func getInputChan() <-chan int {

	input := make(chan int, 100)

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for num := range numbers {
			input <- num
		}

		close(input)
	}()

	return input
}

func getSquareChan(input <-chan int) <-chan int {
	output := make(chan int, 100)

	go func() {
		for num := range input {
			output <- num * num
		}

		close(output)
	}()

	return output
}

func merge(outputsChan ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	merged := make(chan int, 100)

	wg.Add(len(outputsChan))

	output := func(sc <-chan int) {
		for sqr := range sc {
			merged <- sqr
		}
		wg.Done()
	}

	for _, optChan := range outputsChan {
		go output(optChan)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
