package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) SersveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello from go web server.</h1>")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
