package main

import "fmt"

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	intSlice := []int32{4, 5, 6}
	fmt.Println("The length is", len(intSlice), "and the capacity is", cap(intSlice))
	intSlice = append(intSlice, 7) // re-allocate memory to increase the capacity
	fmt.Println(intSlice)
	fmt.Println("The length is", len(intSlice), "and the capacity is", cap(intSlice))

	intSlice2 := []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Println("The length is", len(intSlice), "and the capacity is", cap(intSlice))

	intSlice3 := make([]int32, 5, 20)
	fmt.Println(intSlice3)
	fmt.Println("The length is", len(intSlice3), "and the capacity is", cap(intSlice3))

	myMap := make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(myMap2["one"])
	age, ok := myMap2["four"]
	if !ok {
		fmt.Println("The key does not exist.")
	} else {
		fmt.Println(age)
	}

	for key, value := range myMap2 {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	for i, v := range intArr {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}