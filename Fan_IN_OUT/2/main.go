package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	timer := time.After(time.Millisecond)
OUT:
	for {
		select {
		case <-timer:
			break OUT
		default:
			run()
		}
	}

	time.Sleep(15 * time.Millisecond)
}

func run() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		for _, msg := range []string{"create user", "update user", "select user"} {
			ch1 <- msg
		}
		close(ch1)
	}()

	go func() {
		for _, msg := range []string{"create order", "update order", "select order"} {
			ch2 <- msg
		}
		close(ch2)
	}()

	go func() {
		for _, msg := range []string{"create task", "update task", "run task"} {
			ch3 <- msg
		}
		close(ch3)
	}()

	output := merge(ch1, ch2, ch3)
	go func() {
		for {
			select {
			case v, ok := <-output:
				if !ok {
					return
				}
				fmt.Println(v)
			default:
				fmt.Println("waiting...")
			}
		}
	}()
}

func merge(sources ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	result := make(chan string)

	output := func(ch <-chan string) {
		for v := range ch {
			result <- v
		}
		wg.Done()
	}

	for _, source := range sources {
		wg.Add(1)
		go output(source)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
