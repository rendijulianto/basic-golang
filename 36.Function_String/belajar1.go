package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExists = strings.Contains("John wick", "wick")
	fmt.Println(isExists)

	//HasPrefix
	//Digunakan untuk deteksi apakah sebuah string (
	//parameter pertama) diawali string tertentu (parameter kedua).
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	//HasSuffix
	//Digunakan untuk deteksi apakah sebuah string (parameter pertama)
	//diakhiri string tertentu (parameter kedua).
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	//Count
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	//Index
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2
	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4

	//Replace

	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	//Repeat
	var str = strings.Repeat("na", 4)
	fmt.Println(str) // "nananana"

	//Split
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	//JOIN
	var data = []string{"banana", "papaya", "tomato"}
	var strJoin = strings.Join(data, "-")
	fmt.Println(strJoin) // "banana-papaya-tomato"

	//Fungsi strings.ToLower()
	var strLower = strings.ToLower("aLAy")
	fmt.Println(strLower) // "alay"
	//Fungsi strings.ToUpper()
	var strToUpper = strings.ToUpper("aLAy")
	fmt.Println(strToUpper) // "alay"
}
