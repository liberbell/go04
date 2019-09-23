package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intsum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intsum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatsum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatsum)

	var b1, b2, b3, bigsum big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(65.2)
	b3.SetFloat64(76.3)

	bigsum.Add(&b1, &b2).Add(&bigsum, &b3)
	fmt.Printf("Bigsum= %.10g\n", &bigsum)

	circleRadius := 15.5
	circumference := circleRadius * math.Pi
	fmt.Printf("Circumference is: %.2f\n", circumference)
}
