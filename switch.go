package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""
	switch dow {
	case condition:
		result := "It's Sunday!"
	}
}
