package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	numWorker := 5

	jobs := make(chan func())

	// wg := &sync.WaitGroup{}
	var wg sync.WaitGroup

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workerFunc(jobs)
		}()
	}

	for i := 0; i < 10; i++ {
		jobs <- someWork(i)
	}

	close(jobs)

	wg.Wait()
	fmt.Println("Done")
}

func workerFunc(jobs <-chan func()) {

	for f := range jobs { // при range канал выполняется до тех пор, пока канал не будет закрыт (без явной обработки)
		f()
	}
}

func someWork(i int) func() {

	return func() {
		fmt.Println("work:", i)
		n := rand.Intn(5)
		time.Sleep(time.Duration(n) * time.Second)
	}

}
