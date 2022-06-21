package main

import "fmt"

// ананоимные функции используются с go -routine

// замыкание это функция, которая ссылается на независимые или свободные состояния.
// Функция запоминает состояние в котором она была инициализирована
func main() {
	/// переменная может присваивать себе функцию в ней лежит функция замыкания, которая запоминает значение с прошлого вызова.
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	// func() {

	// 	fmt.Println("anonymous funnction")
	// }()
}

func increment() func() int {
	count := 5
	return func() int {
		count++
		// возвращаем функцию, а не вызываем  и на примере инкрементации в первый раз она будет содержать +1 во второй +2 и так далее.
		return count
	}
}
