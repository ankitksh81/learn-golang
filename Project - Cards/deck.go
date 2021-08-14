package main

import "fmt"

// Creat new type 'deck' which is a slice of strings
type deck []string // similar to structure in c++

// print() function is a receiver function which receives a argument d
// of type deck.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
