package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "muhammad,ravi,mega\n" +
		"budi,sudarsodo,co\n" +
		"eko,kurniawan,khaneddy"

	reader := csv.NewReader(strings.NewReader(csvString))
	fmt.Println(reader)

	/*
		looping disini kita menggunakan perulangan tanpa henti,
		tapi tenang kita pasang break untuk menghentikannya ketika
		kondisi sudah terpenuhi.
	*/
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
