package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 10
	var donominator int = 3
	result, remainder, error := intDivision(numerator, donominator)
	switch {
	case error != nil:
		fmt.Println(error)
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v, and the remainder is %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division is exact.")
	case 1,2:
		fmt.Printf("The division is close.")
	default:
		fmt.Printf("The division is not close.")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}


func intDivision(numerator int, donominator int) (int, int, error) {
	var error error
	if donominator == 0 {
		error = errors.New("donominator is 0")
		return 0, 0, error
	}
	return numerator / donominator, numerator % donominator, error
}