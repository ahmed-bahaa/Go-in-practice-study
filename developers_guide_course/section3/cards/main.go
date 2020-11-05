package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	cards := newDeck()
	fmt.Println(cards.toString())
	// cards = append(cards, newCard2()) // note to work you need to run both of files $ go run main.go newcCard2.go
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
