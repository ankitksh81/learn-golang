package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my deck.txt")
	// cards := newDeckFromFile("my deck.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
