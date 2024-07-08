package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

/*
1. newDeck

2. Print Deck

3. Shuffle the Deck

4. Deal of the cards

5. Save the deck To File

6. New Deck from saved file
*/

type Deck []string

func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) PrintDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d Deck) Shuffle() {
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

func DealDeck(d Deck, handshake int) (Deck, Deck) {
	return d[:handshake], d[handshake:]
}

func (d Deck) ToString() string {
	return strings.Join(d, ",")
}
func (d Deck) SaveToFile(filename string) error {
	err := os.WriteFile(filename, []byte(d.ToString()), 0600)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("File saved to", filename)
	return nil
}

func NewDeckFromFile(filename string) (Deck, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading deck.txt\nError:", err)
		os.Exit(1)
	}

	cardstring := strings.Split(string(bs), ",")

	return cardstring, err
}
