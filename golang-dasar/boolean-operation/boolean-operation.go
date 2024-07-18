package main

import "fmt"

func main(){
	var nilaiUTS = 90
	var nilaiUAS = 80

	var isUTSLulus bool = nilaiUTS > 80
	var isUASLulus bool = nilaiUAS >= 80

	var isMahasisLulus bool = isUTSLulus && isUASLulus
	fmt.Println("Is Mahasiswa Lulus: ", isMahasisLulus)
}