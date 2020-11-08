package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// not a good way in declaration
	//alex := person{"ahmed", "bahaa"}

	alex := person{firstName: "ahmed", lastName: "Bahaa"}
	fmt.Println(alex)
}
