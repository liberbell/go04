package main

import "fmt"

func main() {
	var p *int

	if p != nil {
		fmt.Println("Value of P: ", *p)
	} else {
		fmt.Println("P is nil.")
	}
}
