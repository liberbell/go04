package main

import "fmt"

var isConnected bool = false

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

func connect() {
	isConnected = true
	fmt.Println("Connected to database.")
}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnect from database.")
}
