package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "jim",
		contact: contactInfo{
			email:   "alex_jim@gmail.com",
			zipCode: 555,
		},
	}

	alexPointer := &alex
	alexPointer.updateFirstName("btata")
	alex.print()
}

func (pointerToPerson *person) updateFirstName(name string) {
	(*pointerToPerson).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func main() {
// 	// not a good way in declaration
// 	//alex := person{"ahmed", "bahaa"}

// 	alex := person{firstName: "ahmed", lastName: "Bahaa"}
// 	fmt.Println(alex)

// 	// third way in declaration
// 	var person2 person
// 	fmt.Println(person2)
// 	fmt.Printf("%+v", person2)

// 	fmt.Println("")

// 	person2.firstName = "blaf"
// 	person2.lastName = "blalast"
// 	fmt.Printf("%+v", person2)
// }
