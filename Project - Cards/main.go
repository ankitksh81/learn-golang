package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()} //dynamic array
	cards = append(cards, "Six of Spades")

	var i int
	var card string

	for i, card = range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
