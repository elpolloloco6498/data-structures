package main

import "fmt"

func main() {
	str := "hello"
	fmt.Println(str)
	numbers := []int{1,4,6,3}
	
	fmt.Println(numbers)
	numbers = append(numbers, 44)
	fmt.Println(numbers)
}