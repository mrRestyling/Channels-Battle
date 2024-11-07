package main

import (
	"fmt"
	"math/rand"
	"time"
)

//     Написать генератор случайных чисел

// Асинхронное взаимодействие в Go
// Небуфферезированный канал.
// Будем асинхронно писать туда случайные числа и закроем его, когда закончим писать.

func randNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(10)
		}
		close(out)
	}()
	return out
}

func main() {
	for num := range randNumsGenerator(10) {
		fmt.Println(num)
	}
}
