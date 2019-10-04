package main

import "fmt"

func main() {
	defer fmt.Println("Close the file.")
	fmt.Println("Open the file.")

	defer fmt.Println("Statement 1.")
	defer fmt.Println("Statement 2.")
	defer fmt.Println("Statement 3.")
	defer fmt.Println("Statement 4.")
	fmt.Println("Undeferd statement.")
}

func myfunc() {
	defer fmt.println("Deffered in the function")
	fmt.Println("Not deffered in the function")
}
