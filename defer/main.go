package main

import "fmt"

func main() {
	a := mul()
	fmt.Println("7. ", a)
}

func mul() int {
	fmt.Println("11. Starting program")
	defer fmt.Println("defer 12. Ending program")
	defer fmt.Println("defer 13. Ending program")
	defer fmt.Println("defer 14. Ending program")

	a, b := 10, 99
	c := a / b

	fmt.Printf("19. %d * %d = %d\n", a, b, c)
	return c
}
