package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	// Variadic argument dari number biasa
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println("To9/tal: ", total)

	// Variadic argument dari slice
	numbers := []int{10, 10, 10, 10}
	totalSlice := sumAll(numbers...)
	fmt.Println("Total Slice: ", totalSlice)
}
