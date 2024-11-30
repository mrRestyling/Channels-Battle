package main

import "fmt"

func main() {

	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)
	numsCh <- []int{10, 10, 10}

	res := <-sumCh // 30

	fmt.Println(res)
}

// SumWorker - суммирует переданные числа из канала numsCh и передает результат в канал sumC
func SumWorker(numsCh chan []int, sumCh chan int) {

	for arr := range numsCh {
		sumCh <- Sum(arr)
	}

}

// Sum - вспомогательная функция, которая суммирует числа массива
func Sum(arr []int) int {

	result := 0

	for _, num := range arr {
		result += num
	}

	return result
}
