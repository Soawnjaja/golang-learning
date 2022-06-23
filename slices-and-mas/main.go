package main

import (
	"errors"
	"fmt"
)

// ----Слайсы

func main() {
	// инициализировали слайс и сразу передали значение.
	//	messages := []string{"1", "2", "3"}

	// инициализируем слайс с помощью Make(), указываем количество елементов - 5
	messages := make([]string, 5)
	//  функция append  принимает слайс и елементы которые аппендим .
	//  при аппенде капасити возрастает в большем количестве в зависимости от размера слайса
	// переалокация до опредленных значение массив(слайс) удваивается в размерах. не дешевая операция
	messages = append(messages, "6")
	messages = append(messages, "7")
	messages = append(messages, "8")
	messages = append(messages, "9")
	messages = append(messages, "10")
	messages = append(messages, "11")
	printMessage(messages)
	fmt.Println(messages)
	// капасити равно 20.
	fmt.Println(cap(messages))

}

// Slice хранит ссылку на массив и когда мы передаем слайс в функцию, мы передаем ссылку а значит можем менять значения внутри слайса внутри функции.
func printMessage(messages []string) error {
	if len(messages) == 0 {
		return errors.New(" slice is empty")
	}
	messages[0] = "4"
	fmt.Println(messages)
	return nil
}

// ---- Массив
// func main() {
// 	// в массивах указываем заранее количество елементов , которое там будет жить [3]string{} - массив может быть создан пустым.
// 	messages := [3]string{"4", "9", "12"}
// 	// перезаписать по индексу как и в js
// 	messages[1] = "5"
// 	fmt.Println(messages)

// 	/*	fmt.Println(messages[3]) нет доступа к 4 елементу так как массив фиксирован, так же мы не можем передавать массив в аргумент
// 		(messages [3]string) , если он содержит больше елементов.
// 	*/
// }
// // передавая массив в аргументы функции мы работаем с копией, поэтому при попытке изменить значение внутри массива изменений не произойдет
// func printMessage(messages [3]string) error {
// 	if len(messages) == 0 {
// 		return errors.New("empty array")
// 	}
// 	messages[2] = "10"
// 	fmt.Println(messages)
// 	return nil
// }

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
