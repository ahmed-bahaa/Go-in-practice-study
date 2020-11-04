package main

func main() {
	// var card string = "Ace of spades"
	cards := newDeck()
	// cards = append(cards, newCard2()) // note to work you need to run both of files $ go run main.go newcCard2.go

	cards.print()
}
