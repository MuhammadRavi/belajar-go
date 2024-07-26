package main

import (
	"fmt"
	"os"
)

func main() {
	OS := os.Args
	fmt.Println(OS)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err)
	}
}
