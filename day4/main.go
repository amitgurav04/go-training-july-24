package main

import "fmt"

func main() {

	myslice := [6]string{"AA", "BB", "CC", "DD", "EE", "FF"}

	//fmt.Println("Before:", myslice)

	//updateSlice(myslice)

	for _, v := range myslice {
		fmt.Println(v)

	}

	_, y := myFunction()

	fmt.Println(y)

	//fmt.Println("After: ", myslice)
}

func myFunction() (a, b int) {
	return 1, 2
}

func updateSlice(s [6]string) {
	s[0] = "AAA"
}
