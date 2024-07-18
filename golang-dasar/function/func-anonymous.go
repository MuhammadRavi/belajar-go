package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// function anonymous: function yang tidak diberikan nama.

	// Ada beberapa cara dalam pembuatan function anonymous
	// diantaranya:

	// 1. menampung function anonymous kedalam variable.
	blackList := func(name string) bool {
		return name == "Anjing"
	}
	fmt.Println("================ Cara Pertama ===================")
	registerUser("Anjing", blackList)

	// 2. membuat function anonymous pada saat proses pemanggilan
	// function utamanya.
	fmt.Println("\n================ Cara Kedua ===================")
	registerUser("Muhammad Ravi Mega Arasy", func(name string) bool {
		return name == "Anjing"
	})

}
