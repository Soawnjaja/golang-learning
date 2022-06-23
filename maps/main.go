package main

import "fmt"

// мапа не упорядочный массив данных ключ + значение как объект в JS.
// ключ не может дублироваться.
// func main() {
// 	users := map[string]int{
// 		"Vlad":   15,
// 		"Kostya": 48,
// 		"Maxim":  32,
// 	}
// вывод по ключу - возвращаем значение
// fmt.Println(users["Vlad"])
// если ключа нет то выводит 0 значение
// можно присваивать 2 значния по ключу к примеру хотим получить значние и есть ли вообще оно exists или ok
// 	age, exists := users["Kostya"] // if ["Serega"] = else
// 	if exists {
// 		fmt.Println("Kostya", age)
// 	} else {
// 		fmt.Println("Kostya нет в списке")
// 	}

// }
// итерация по мапе 2 значение key и value соответсвенно в цикле range.
func main() {
	users := map[string]int{
		"Vlad":   15,
		"Kostya": 48,
		"Maxim":  32,
	}
	// инициализация мапы нужна что бы не словить панику
	// someusers := make(map[string]int ) строкове значение ключ и числовое знаничение так же инициализируется через make(метод[ключ]значение)

	// удаление по ключу - delete(users,"Vlad")
	users["Karina"] = 18
	for key, value := range users {
		fmt.Println(key, value)
	}
	// так же принима встроенная функция len() узнать количество елементов в мапе

}
