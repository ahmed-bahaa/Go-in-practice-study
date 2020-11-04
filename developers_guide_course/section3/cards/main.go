package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of spades"
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, newCard2()) // note to work you need to run both of files $ go run main.go newcCard2.go

	for i, card := range cards {
		fmt.Println(i, card)
	}

	// fmt.Println(card)
	// fmt.Println(newCard2())
}

func newCard() string {
	return "Five of Diamonds"
}
