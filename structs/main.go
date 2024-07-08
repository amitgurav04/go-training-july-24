package main

import "fmt"

type state struct {
	stateName string
	country   string
}
type contactInfo struct {
	email   string
	zipCode int
	state
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(correctName string) {
	p.firstName = correctName
	//fmt.Println("========Inside UpdateName==============")
	//p.print()
	//fmt.Println("======End Inside UpdateName============")
}

func mai1() {
	person1 := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "john.doe@gmail.com",
			zipCode: 94000,
			state: state{
				stateName: "Maharashtra",
				country:   "India",
			},
		},
	}
	person1.print()

}

func main() {
	p := person{firstName: "Prakharr", lastName: "Sharma"}

	fmt.Println("Before Correction")
	p.print()

	pointerP := &p

	(*pointerP).updateName("Prakhar")
	fmt.Println("After Correction")
	p.print()

}
