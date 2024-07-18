package main

import "fmt"

func main() {
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Muhammad"
	newSlice[1] = "Ravi"

	fmt.Println("Data Slice: ", newSlice)
	fmt.Println("Capacity Slice: ", cap(newSlice))
	fmt.Println("Length Slice: ", len(newSlice))

	fmt.Println(" ============================ ")
	newSlice2 := append(newSlice, "Mega")
	fmt.Println("Data Slice: ", newSlice)
	fmt.Println("Data Slice 2: ", newSlice2)
	fmt.Println("Capacity Slice: ", cap(newSlice2))
	fmt.Println("Length Slice: ", len(newSlice2))

	fmt.Println(" ========= Modify =========== ")
	newSlice2[0] = "Budi"
	fmt.Println("Data Slice: ", newSlice)
	fmt.Println("Data Slice 2: ", newSlice2)
}
