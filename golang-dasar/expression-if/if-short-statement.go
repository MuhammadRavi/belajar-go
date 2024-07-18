package main

import "fmt"

func main() {
	name := "Muhammad"

	// 1. Short statement: kita membuat variable langsung dalam if'nya
	//    seperti contoh dibawah yaitu: "length := len(name);".
	// 2. ketika kita menggunakan short statement maka kita harus akhiri dengan ";"
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
