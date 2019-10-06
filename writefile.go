package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from go."
	file, err := os.Create("./fromstring.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Add done with file of %v characters\n", ln)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
