package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}

	dCh := make(chan struct{})
	defer close(dCh)

	output := multiply(dCh, add(dCh, generator(dCh, input)))

	for result := range output {
		fmt.Println(result)
	}
}
func multiply(d chan struct{}, i chan int) chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for data := range i {
			result := data * 2

			select {
			case <-d:
				return
			case out <- result:
			}
		}

	}()

	return out
}
func add(d chan struct{}, i chan int) chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for data := range i {
			result := data + 1

			select {
			case <-d:
				return
			case out <- result:

			}
		}
	}()
	return out
}

func generator(d chan struct{}, input []int) chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for _, data := range input {
			select {
			case <-d:
				return
			case out <- data:
			}
		}
	}()
	return out
}
