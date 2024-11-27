package main

// 1
// go build -gcflags=-m
// Сообщения «escapes to heap» нет, что говорит об отсутствии перемещения переменной из стека в кучу.

// func main() {
// 	var a int64 = 10

// 	print(a)
// }

// 2
// Переменная может попасть в кучу из‑за объёма занимаемой памяти
// func main() {
// 	s1 := make([]int, 10)
// 	s2 := make([]int, 10000)

// 	print(s1[0])
// 	print(s2[0])
// }
