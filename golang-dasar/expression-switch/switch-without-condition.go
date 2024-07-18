package main

import "fmt"

func main() {
	name := "Muhammad"
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama Terlalu Panjang")
	case length < 5:
		fmt.Println("Nama Sudah Benar")
	}
}
