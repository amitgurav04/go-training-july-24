package main

import "fmt"

func abc() {

}

func main() {

	//var b int
	b := "abc"
	//c := 55

	fmt.Println(b)
	//fmt.Println(c)

	/*	cards := []string{
			"Ace Of Spades",
			"Two Of Spades",
			"Three Of Spades",
			"Four Of Spades",
		}
		fmt.Println("Before")
		fmt.Println("Length:", len(cards))
		fmt.Println("Capacity:", cap(cards))

		cards = append(cards, "Five Of Spades")

		fmt.Println(cards)
		fmt.Println("After Add")
		fmt.Println("Length:", len(cards))
		fmt.Println("Capacity:", cap(cards))
		cards1 := cards[0:1]
		cards2 := cards[2:]
		cards1 = append(cards1, cards2...)

		fmt.Println(cards1)
		fmt.Println("After Delete")
		fmt.Println("Length:", len(cards))
		fmt.Println("Capacity:", cap(cards))

	*/ /*	fmt.Println("Before Update")
		printSlice(cards)

		updateSlice(cards)

		fmt.Println("After Update")
		printSlice(cards)
	*/
}

func updateSlice(a []string) {

	a[1] = "TWO Of Spades"
	fmt.Println("\n\n------------Cards inside update function-----------")
	fmt.Println(a)
	fmt.Println("\n\n------------End Cards inside update function-----------")

}

func printSlice(a []string) {

	fmt.Println(a)

}
