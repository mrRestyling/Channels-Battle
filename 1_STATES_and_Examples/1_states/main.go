package main

import (
	"fmt"
	"sync"
)

func main() {

	// nilClose()
	// nilRead()
	// nilWrite()

	// closedCLOSE()
	closedREAD()
	// closedWRITE()
}

func closedCLOSE() {
	var wg sync.WaitGroup

	ch := make(chan int)
	defer close(ch)

	wg.Add(1)

	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()

	ch <- 1
	wg.Wait()

	close(ch)
}

func closedREAD() {

	wg := &sync.WaitGroup{}

	ch := make(chan int)

	close(ch)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	wg.Wait()
}

func closedWRITE() {

	ch := make(chan int)
	close(ch)

	ch <- 1
}

func nilClose() {

	var ch chan int

	close(ch)
}

func nilRead() {

	var ch chan int

	<-ch
}

func nilWrite() {

	var ch chan int

	ch <- 1
}
