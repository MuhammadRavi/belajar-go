package main

import "fmt"

func main() {
	name := "ravi"

	switch name {
	case "ravi":
		fmt.Println("Hello ravi")
	case "joko":
		fmt.Println("Hello joko")
	default:
		fmt.Println("Hallo, Boleh Kenalan ? ")
	}
}
