package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 1
	var result = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 5
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString(myString))
	fmt.Println(len("ガンマ"))
	fmt.Println(utf8.RuneCountInString("ガンマ"))

	var myRune rune = 'a'
	fmt.Println(myRune)
}