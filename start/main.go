package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	message, error := enterTheClub(56)
	if error != nil {
		// логирование ошибки log.Fatal(error) передаем ошибку и выходим из приложения с фиксацией времени.
		log.Fatal(error)
		// просто выходим из фукнции не передавая значения в return
		return
	}
	fmt.Println(message)
}

// принимае множественные значения строка + ошибка
func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		// если ошибка равна nil, то ошибки нет.
		return "Please enter, but be safe", nil
	} else if age >= 56 {
		// в случае если мы словили ошибку обозначаем ее конструкцией новой ошибки.
		return "U are so old dude", errors.New("you are to old")
	} else if age >= 18 && age <= 45 {
		return "U are really need to this music?", nil
	}
	// конструкция ошибки errors.New("low register on eng lang")
	return "U cant enter, come again when grow up.", errors.New("you are to yang man")
}
