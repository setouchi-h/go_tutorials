package main

import "fmt"

func main() {
	p := new(int32)
	var i int32
	fmt.Println("The value p is", p)
	fmt.Println("The value p point to is", *p)
	fmt.Println("The value i is", i)
	p = &i
	*p = 1
	fmt.Println("The value p is", p)
	fmt.Println("The value p point to is", *p)
	fmt.Println("The value i is", i)

	slice := []int32{1,2,3}
	sliceCopy := slice // メモリアドレスはsliceと同じ(メモリアドレスをコピーしているだけ)
	fmt.Println("The value slice is", slice)
	fmt.Println("The value sliceCopy is", sliceCopy)
	sliceCopy[0] = 100
	fmt.Println("The value slice is", slice)
	fmt.Println("The value sliceCopy is", sliceCopy)

	thing1 := [5]float64{1,2,3,4,5}
	fmt.Printf("The memory location of thing1 array is %p\n", &thing1)
	result1 := square(thing1)
	fmt.Printf("The memory location of result1 array is %p\n", &result1)

	result2 := square2(&thing1)
	fmt.Printf("The memory location of result2 array is %p\n", &result2)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 array is %p\n", &thing2) // 関数に引数を渡す際に、配列の値がコピーされる
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func square2(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 array is %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}