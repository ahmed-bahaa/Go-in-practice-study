package main

func main() {
	// var card string = "Ace of spades"
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, newCard2()) // note to work you need to run both of files $ go run main.go newcCard2.go

	cards.print()
	// fmt.Println(card)
	// fmt.Println(newCard2())
}

func newCard() string {
	return "Five of Diamonds"
}
