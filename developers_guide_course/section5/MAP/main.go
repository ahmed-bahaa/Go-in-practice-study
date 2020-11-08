package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":  "#fff222",
		"blue": "#fff6564564",
	}

	//declare an empty map 1
	// var colors2 map[string]string

	//decalre an empty map2
	color2 := make(map[string]string)
	fmt.Println("Printing empty colors2: ", color2)

	color2["test-color"] = "#fd5555"
	fmt.Println("Printing non empty colors2: ", color2)

	delete(color2, "test-color")
	fmt.Println("Printing empty colors2 after delete: ", color2)

	fmt.Println(colors)
}
