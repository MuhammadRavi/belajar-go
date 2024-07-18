package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	nilai, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Nilai", nilai)
	} else {
		fmt.Println("Error", err.Error())
	}
}
