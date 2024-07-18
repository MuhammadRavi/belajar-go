package main

import "fmt"

func main(){
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println("Konversi Nilai int 32: ",nilai32)
	fmt.Println("Konversi Nilai int 64: ",nilai64)
	fmt.Println("Konversi Nilai int 16: ",nilai16)
}