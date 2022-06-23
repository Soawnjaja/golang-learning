package main

import "fmt"

// цикл для массива и слайса range
func main() {
	messages := []string{
		"message0",
		"message1",
		"message2",
		"message3",
		"message4",
	}
	// можно заменять
	for _, messages := range messages {
		fmt.Println(messages)
	}
}

// конструкция остановки цикла
// 	counter := 0
// 	for  {
// 		if counter = 100 {
// 			break
// 		}
// 		counter ++
// 	}
// }

//  матрица представляет собой слайс и каждая ячейка слайса храните в себе слайс елементов.

// func main() {

// 	//  инициализируем матрицу в мейк
// 	// цикл работает пока не заполнит 10 елементов внутри каждой
// 	matrix := make([][]int, 10)
// 	for x := 0; x < 10; x++ {

// 		for y := 0; y < 10; y++ {
// 			matrix[y] = make([]int, 10)
// 			matrix[y][x] = x
// 		}
// 		fmt.Println(matrix[x])
// [0 0 0 0 0 0 0 0 0 0]
// [0 1 0 0 0 0 0 0 0 0]
// [0 0 2 0 0 0 0 0 0 0]
// [0 0 0 3 0 0 0 0 0 0]
// [0 0 0 0 4 0 0 0 0 0]
// [0 0 0 0 0 5 0 0 0 0]
// [0 0 0 0 0 0 6 0 0 0]
// [0 0 0 0 0 0 0 7 0 0]
// [0 0 0 0 0 0 0 0 8 0]
// [0 0 0 0 0 0 0 0 0 9]
// 	}

// }
