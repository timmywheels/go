package main

func main() {
	cards := newDeckFromFile("cards.txt")

	//cards.print()

	cards.shuffle()

	cards.print()

	cards.saveToFile("cards.txt")

	//fmt.Println(cards.toString())
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
}
