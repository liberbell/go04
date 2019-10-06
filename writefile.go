package main

import (
	"io"
	"os"
)

func main() {
	connect := "Hello from go."
	file, err := os.Create("./fromstring.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)
}

func checkError() {
	if err != nil {
		panic(err)
	}
}
