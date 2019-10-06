package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "./hello.txt"

	content, err := ioutil.ReadFile(filename)
	checkError(err)

	fmt.Println("Read from file: ", content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
