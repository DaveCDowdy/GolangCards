package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	hand, _ := deal(cards, 3)
	hand.print()
}
