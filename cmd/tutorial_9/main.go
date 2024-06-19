package main

import "fmt"

// dead lockエラーの発生
// func main() {
// 	c := make(chan int)
// 	c <- 1
// 	i := <- c
// 	fmt.Println(i)
// }

func main() {
	c := make(chan int)
	go process(c)
	for v := range c {
		fmt.Println(v)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}