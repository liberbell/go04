package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response type: %T\n", resp)
}
