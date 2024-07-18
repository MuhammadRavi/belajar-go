package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		if counter == 5 {
			break
		}
		fmt.Println("Perulangan ", counter)
	}

}
