package main

import "net/http"

func main() {
	url := "https://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
}
