package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)
	fmt.Println(numbers[2])

	fmt.Println("The number is colors: ", len(colors))
	fmt.Println("The number is numbers: ", len(numbers))
}
