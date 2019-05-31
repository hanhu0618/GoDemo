package main

func main() {
	//var card string = "Ace of Spades"
	// card := newCard()
	// cards := deck{"Ace of Dimands", newCard()}
	//cards = append(cards, "Six of Spades") //add elements
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// word := ([]byte)("Hi there!")
	// fmt.Println(word)

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
