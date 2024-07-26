package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"muhammad", "ravi", "mega"})
	writer.Write([]string{"budi", "ravi", "mega"})
	writer.Write([]string{"eko", "ravi", "mega"})

	writer.Flush()
}
