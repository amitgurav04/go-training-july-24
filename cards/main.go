package main

import (
	"fmt"
	"github.com/amitgurav04/cards/deck"
)

func main() {
	d := deck.NewDeck()

	//fmt.Println("====================Before Shuffle=============")
	d.PrintDeck()
	d.Shuffle()
	//fmt.Println("====================After Shuffle==============")
	//d.PrintDeck()

	fmt.Println("\n======================Deck in String Format========================")
	fmt.Println(d.ToString())
	d.SaveToFile("deck.txt")

	fmt.Println("\n======================Read Deck from file========================")
	d2, err := deck.NewDeckFromFile("deck1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("=============================Deck after reading file=============================")
	d2.PrintDeck()

	/*	fmt.Println("\n====================Before Deal=============")
		d.PrintDeck()

		d1, d2 := deck.DealDeck(d, rand.Intn(52))

		fmt.Println("\n====================After Deal==============")
		fmt.Println("\nDeck 1:")
		d1.PrintDeck()
		fmt.Println("\n====================End Deck 1==============")

		fmt.Println("\n=====================Deck 2:================")
		d2.PrintDeck()
		fmt.Println("\n====================End Deck 2==============")
	*/
}
