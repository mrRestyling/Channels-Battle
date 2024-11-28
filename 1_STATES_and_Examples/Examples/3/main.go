package main

import "fmt"

func main() {
	run()
}

func run() {

	ch := make(chan string)

	go func() {
		for n := range ch {
			fmt.Println(n)
		}
	}()

	ch <- "1"
	ch <- "2"
}
