package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)
}
