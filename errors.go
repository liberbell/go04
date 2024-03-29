package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.txt")
	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myerror := errors.New("My error message.")
	fmt.Println(myerror)

	attendence := map[string]bool{
		"Ann":  true,
		"Mike": true}

	attended, ok := attendence["M"]
	if ok {
		fmt.Println("Mike attended?", attended)
	} else {
		fmt.Println("No info for Mike.")
	}
}
