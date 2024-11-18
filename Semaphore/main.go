package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	numWorker := 3

	sem := make(chan struct{}, numWorker)

	defer close(sem)

	var wg sync.WaitGroup

	for i := 0; i < 15; i++ {

		wg.Add(1)

		sem <- struct{}{}
		// log.Println("1")

		go func() {
			// log.Println("2")
			defer wg.Done()
			someWork(i)
			<-sem

		}()
	}
	wg.Wait()

	fmt.Println("Final")

}

func someWork(i int) {
	// log.Println("3")
	n := rand.Intn(5)
	fmt.Println("run job: ", i, n)
	time.Sleep(time.Duration(n) * time.Second)

}
