package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()} //dynamic array
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
