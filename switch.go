package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)
}
