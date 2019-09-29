package main

import "fmt"

func main() {
	dosomething()
}

func dosomething() {
	fmt.Println("Doing something.")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}
