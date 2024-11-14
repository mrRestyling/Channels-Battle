package main

import "fmt"

// func main() {

// 	ch := make(chan int, 1)

// 	go func() {

// 		<-ch
// 		fmt.Println("go")

// 	}()

// 	ch <- 100
// 	fmt.Println("main")

// }

func main() {

	noBuf()
	Buf()

}

func noBuf() {

	ch := make(chan int)

	go func() {

		<-ch
		fmt.Println("go")

	}()

	ch <- 100
	fmt.Println("main")

}

func Buf() {

	ch := make(chan int, 1)

	go func() {

		<-ch
		fmt.Println("go")

	}()

	ch <- 100
	fmt.Println("main")

}
