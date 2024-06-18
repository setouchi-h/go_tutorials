package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "résumé"
	indexed := myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe lenght of 'myString' is %v\n", len(myString))

	myString2 := []rune(myString)
	indexed2 := myString2[1]
	fmt.Printf("%v, %T\n", indexed2, indexed2)
	for i, v := range myString2 {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe lenght of 'myString2' is %v\n", len(myString2))

	myRune := 'a'
	fmt.Printf("\nmyRune = %v, %T\n", myRune, myRune)

	strSlice := []string{"t", "e", "s", "t"}
	fmt.Printf("\nstrSlice = %v, %T\n", strSlice, strSlice)
	catStr := ""
	for i := range strSlice {
		catStr += strSlice[i] // 毎回変数を新規に作るので遅い
	}
	fmt.Println(catStr)

	strSlice2 := []string{"t", "e", "s", "t"}
	var strBuilder strings.Builder
	for i := range strSlice2 {
		strBuilder.WriteString(strSlice2[i]) // ここはinternalにする
	}
	catStr2 := strBuilder.String() // 最後にここで変数を新規作成するので速い
	fmt.Println(catStr2)
}