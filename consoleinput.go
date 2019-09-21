package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Printf("Hello World!\n")
	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
}
