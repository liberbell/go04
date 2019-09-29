package main

import "fmt"

func main() {
	dosomething()
	addValues(1, 2)
	fmt.Println(addValues)
}

func dosomething() {
	fmt.Println("Doing something.")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}
