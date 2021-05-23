package main

import "fmt"

func main() {

	cards := newDeck()
	cards.saveToFile("my_cards")

	fmt.Println(cards.toString())
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
