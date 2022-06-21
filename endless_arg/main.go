package main

import "fmt"

// ищем в списке минимальное значение
func main() {
	fmt.Println(findMin(44, 33, 22, -55, 25044, -300, 532, -1))
}

// не ограниченное количество аргументов в функции
// ищем минимальное число в cлайсе "..." - неограниченное количество
func findMin(numbers ...int) int {
	//встроенная функция len() передаем строку, массив, слайс , канал и  числа = длинна списка
	if len(numbers) == 0 {
		return 0
	}
	// если в списке есть хоть один елемент, инициализируем  переменную  мин и кладем туда массив
	min := numbers[0]
	// итерируемся по списку, сравниаем i c 0 в случае,  если находим min = i значению.
	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}

// конструкция switch (более читабельная, чем if'ы так же как в js не считая break)
// func prediction(dayOfWeek string) string {

// 	switch dayOfWeek {
// 	case "monday":
// 		return "be patient, week just begin"
// 	case "tuesday":
// 		return "fucking thuesday, but feel better"
// 	case "wednesday":
// 		return "this is the middle of the weeeek, oh eeee"
// 	case "thursday":
// 		return "Thursday, check your tusk manager, week it almost end"
// 	case "friday":
// 		return "To laaaate, friday friday friday"
// 	case "saturday":
// 		return "Nevermind whats day is today, i hungry to code"
// 	case "sunday":
// 		return "Last day of week, need to meet with my gf."
// 	default:
// 		return "Please enter day of week"

// 	}
// }
