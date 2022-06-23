package main

import "fmt"

// учимся  паниковать
// func main() {

// 	messages := []string{
// 		"msg1",
// 		"msg2",
// 		"msg3",
// 		"msg4",
// 	}
// 	// паникуем так как по длинна массива 0 1 2 3 и пытаемся назначить 4 индексу значение.
// 	messages[4] = "msg5"
// 	panic(" aaaaaaa Help")
// 	fmt.Println(messages)

// }

/// логика defer  функция выполняется в случае выполнения остальных исполнений в функции. она будет выполнена последней.
// вызов defer имеет over head в 50млсекунд
func main() {
	defer handlerPanic()

	messages := []string{
		"message1",
		"message2",
		"message3",
		"message4",
	}
	messages[4] = "message5"
	fmt.Println(messages)
}

// востанавливаемся из паники в случае ошибки,  messages не выполняются так как имеют ошибку и падают, происходит прекращение функции после чего
// выполняется то что идет в defer.
func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("Recovered from panic")

}
