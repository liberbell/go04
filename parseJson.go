package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	contents := contentFromServer(url)

	fmt.Println(contents)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}
