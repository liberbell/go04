package main

import "fmt"

func main() {
	fmt.Printf("Connection open: %v\n", isConnected)
	doSomething()
	fmt.Printf("Connection open: %v\n", isConnected)
}

func doSomething() {
	connect()
	fmt.Println("Deffering disconnect!")
	defer disconnect()
	fmt.Printf("connection open %v\n", isConnected)
	fmt.Println("Doing something.")
}

func connecct() {
	isConnected = true
	fmt.Println("Connected to database.")
}
