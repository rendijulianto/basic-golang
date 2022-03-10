package main

import (
	"fmt"
	"reflect"
)

func main() {

	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("Tipe Variable :", reflectValue.Type())

	if (reflectValue.Kind() == reflect.Int) {
		fmt.Println("Nilai Variable : ", reflectValue.Int());
	}

}
