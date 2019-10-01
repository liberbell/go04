package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.txt")
	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}
}
