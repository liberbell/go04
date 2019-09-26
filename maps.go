package main

import "fmt"

func main() {
	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "washington"
	states["OR"] = "oregon"
	states["CA"] = "california"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)
}
