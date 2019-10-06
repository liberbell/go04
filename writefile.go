package main

import "os"

func main() {
	connect := "Hello from go."
	file, err := os.Create("./fromstring.txt")
	checkError(err)
	defer file.Close()

}

func checkError() {
	if err != nil {
		panic(err)
	}
}
