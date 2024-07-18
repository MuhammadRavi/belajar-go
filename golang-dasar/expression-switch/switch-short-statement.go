package main

import "fmt"

func main() {
	name := "Muhammad"

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
}
