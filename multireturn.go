package main

import "fmt"

func main() {
	n1, l1 := fullName("Zephard", "Beebrox")
	fmt.Printf("Fullname: %v, number of chars: %v\n", n1, l1)

	n2, l2 := fullName("Aurthur", "Beebrox")
	fmt.Printf("Fullname: %v, number of chars: %v\n", n2, l2)
}

func fullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func fullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length := len(full)
	return
}
