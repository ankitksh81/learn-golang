package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Creat new type 'deck' which is a slice of strings
type deck []string // similar to structure in c++

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // use _ whenever we are declaring a variable but not using it
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print() function is a receiver function which receives a argument d
// of type deck.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// converting deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// type(variable) for type casting
// using ioutil func to save file in system
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}
