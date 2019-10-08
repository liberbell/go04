package main

import (
	"fmt"
	"net/http"
)

type Hello struct {
}

func (h Hello) SersveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello from go web server.</h1>")
}

func main() {
	url
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
