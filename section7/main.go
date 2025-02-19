package main

import "fmt"

type intTransformFn = func(int) int

func main() {
	numbers := []int{2, 4, 3, 9, 5, -2, 7}
	doubled := *transformNumbers(&numbers, double)
	squared := *transformNumbers(&numbers, square)

	fmt.Println("Numbers ", numbers)
	fmt.Println("Doubled ", doubled)
	fmt.Println("Squared ", squared)
}

func transformNumbers(numbers *[]int, transform intTransformFn) *[]int {
	result := make([]int, len(*numbers))

	for i, n := range *numbers {
		result[i] = transform(n)
	}

	return &result
}

func double(n int) int {
	return n * 2
}

func square(n int) int {
	return n * n
}
