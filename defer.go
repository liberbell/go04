package main

import "fmt"

func main() {
	defer fmt.Println("Close the file.")
	fmt.Println("Open the file.")

	defer fmt.Println("Statement 1.")
	defer fmt.Println("Statement 2.")

	myfunc()

	defer fmt.Println("Statement 3.")
	defer fmt.Println("Statement 4.")
	fmt.Println("Undeferd statement.")

	x := 1000
	defer fmt.Println("Value of x: ", x)
	x++
	fmt.Println("Value of after incrementing: ", x)
}

func myfunc() {
	defer fmt.Println("Deffered in the function")
	fmt.Println("Not deffered in the function")
}
