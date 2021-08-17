package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.saveToFile("my deck.txt")
	cards := newDeckFromFile("my deck.txt")
	// fmt.Println([]byte(cards.toString()))

	fmt.Println(cards)
}

// 21 done
