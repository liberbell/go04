package main

import "fmt"

func main() {
	dosomething()
	sum := addValues(1, 2)
	fmt.Println("Sum: ", sum)

	sum = addAllValues(20, 33, 45)
	fmt.Println("Sum: ", sum)
}

func dosomething() {
	fmt.Println("Doing something.")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
