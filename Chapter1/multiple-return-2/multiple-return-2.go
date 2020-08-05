package main

import (
	"fmt"
)

func Names() (first string, second string) {
	first = "go"
	second = "Lovers"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
