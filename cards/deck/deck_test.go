package deck

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d) != 52 {
		t.Error("Expected deck length to be 52, but got ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first card to be Ace of Spades, but got ", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Error("Expected last card to be King of Diamonds, but got ", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	d := NewDeck()

	err := d.SaveToFile("_deckTesting")

	if err != nil {
		fmt.Println(err)
	}

	loadedDeck, err := NewDeckFromFile("_deckTesting")
	if len(loadedDeck) != 52 || err != nil {
		t.Error("Expected deck length to be 52, but got ", len(loadedDeck))
	}

	err = os.Remove("_deckTesting")

	if err != nil {
		fmt.Println(err)
	}
}
