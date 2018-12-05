package main

func main() {
	cards := newDeckFromFile("MyCards")
	//cards.saveToFile("MyCards")
	cards.shuffle()
	cards.print()

	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
}
