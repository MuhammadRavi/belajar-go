package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("========= Parsing String To Boolean =========")
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error:" + err.Error())
	}

	fmt.Println("\n========= Parsing String To Integer =========")
	integer, err := strconv.Atoi("100000")
	if err == nil {
		fmt.Println(integer)
	} else {
		fmt.Println("Error:" + err.Error())
	}

	fmt.Println("\n========= Parsing Integer To String =========")
	var stringInteger string = strconv.Itoa(900000)
	fmt.Println(stringInteger)
}
