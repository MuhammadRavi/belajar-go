package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Ravi"
	names[2] = "Mega Arasy"

	fmt.Println("index 1: ", names[0])
	fmt.Println("index 2: ", names[1])
	fmt.Println("index 3: ", names[2])

	// Static Array
	var values = [3]int{90, 80, 96}
	fmt.Println(values)

	// Dynamic Array, hanya boleh langsung penentuan value
	var dynamicArray = [...]int{
		10,
		20,
		30,
		40,
		50,
	}
	fmt.Println("Dynamic Array: ", dynamicArray)
	fmt.Println("Panjang Array: ", len(dynamicArray))

}
