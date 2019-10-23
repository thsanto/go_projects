package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "opa@gmail.com",
			zipCode: "3408957",
		},
	}

	jimPointer := &jim

	jimPointer.updateName("Jimington")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
