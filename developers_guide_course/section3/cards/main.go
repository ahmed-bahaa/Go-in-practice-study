package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	card := ""
	card = newCard()

	fmt.Println(card)
	fmt.Println(newCard2()) // note to work you need to run both of files $ go run main.go newcCard2.go
}

func newCard() string {
	return "Five of Diamonds"
}
