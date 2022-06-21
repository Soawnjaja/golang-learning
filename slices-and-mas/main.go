package main

import "fmt"

func main() {
	messages := [3]string{}
}

// функция инит имеет свойство срабатывать при инициализации пакета . сначала отработает фукнцаия инит если она есть.

/// указатели
// func main() {
// 	// пустой указатель p = nil не ссылается никуда.
// 	number := 5
// 	var p *int
// 	// cсылаемся на ячейку памяти number
// 	p = &number

// 	//ccылаемся на p в свою очередь она ссылается на ячейку number и number присваевает 10.
// 	p = 10
// 	fmt.Println(p)
// 	message := "How to end this hell and start doing test case"
// 	fmt.Println(message)
// 	changeMessage(&message) //  Референсинг -ссылаемся на *амперсант - передаем область в памяти а не строку.
// 	fmt.Println(message)
// }

// /// во время вызова создается копия переменной и записывается в новый холдер с новым значением
// // Дереферерсинг указываем * по области в памяти обращаемся к значению
// func changeMessage(message *string) { // передаем указатель на переменную, а значит не создаем копии и не занимаем место в оперативе
//  *message += "---from function printMsg" *//
// fmt.Println(message)
// }

// массивы и слайсы
