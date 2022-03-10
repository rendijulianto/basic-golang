package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Nama : ", reflectType.Field(i).Name)
		fmt.Println("Tipe Data : ", reflectType.Field(i).Type)
		fmt.Println("Nilai : ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	//var s1 = &student{Name: "wick", Grade: 2}
	//s1.getPropertyInfo()

	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", s1.Name)
}
