package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello \n")
	writer.WriteString("Selamat Belajar")

	writer.Flush()
}
