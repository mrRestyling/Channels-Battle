package main

import (
	"fmt"
	"sync"
)

// Проверка на RC:     go run -race main.go

var i int

func main() {

	rc()

}

func rc() {

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &mu)
	}

	wg.Wait()

	fmt.Println(i)
}

func worker(wg *sync.WaitGroup, mu *sync.Mutex) {

	defer wg.Done()

	mu.Lock()
	i++
	mu.Unlock()

}
