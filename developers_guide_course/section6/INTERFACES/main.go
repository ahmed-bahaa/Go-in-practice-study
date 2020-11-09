package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	printGreeting(eb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreating())
}

func (englishBot) getGreating() string {
	return "Hi there!"
}

func (spanishBot) getGreating() string {
	return "Hola!"
}
