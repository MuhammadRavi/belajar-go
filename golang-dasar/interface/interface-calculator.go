package main

import "fmt"

// Function Type Declaration
type calculateFunc func(a, b int) int

// Interface Type Declaration
type Calculate interface {
	calculate(a, b int) int
}

func (formatCalculate calculateFunc) calculate(a, b int) int {
	return formatCalculate(a, b)
}

func main() {
	add := calculateFunc(func(a, b int) int {
		return a + b
	})

	fmt.Println("Total:", add.calculate(5, 3))
}
