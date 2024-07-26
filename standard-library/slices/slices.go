package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"josh", "paul", "george", "ringo"}
	value := []int{100, 95, 80, 90}

	fmt.Println("========== Value =============")
	fmt.Println(slices.Max(value))
	fmt.Println(slices.Min(value))

	fmt.Println("\n========== Name =============")
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(names))
}
