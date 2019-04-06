package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
}
