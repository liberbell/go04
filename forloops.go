package main

import "fmt"

func main() {
	sum := 1
	fmt.Println("Sum: ", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 11; i++ {
		sum += i
	}
	fmt.Println("Sum: ", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
}
