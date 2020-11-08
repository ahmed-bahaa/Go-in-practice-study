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

	// third way in declaration
	var person2 person
	fmt.Println(person2)
	fmt.Printf("%+v", person2)

	fmt.Println("")

	person2.firstName = "blaf"
	person2.lastName = "blalast"
	fmt.Printf("%+v", person2)
}
