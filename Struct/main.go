package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	yuri := person{
		firstName: "Yuri",
		lastName:  "Every",
		contactInfo: contactInfo{
			email:   "laduy2208@gmail.com",
			zipCode: 10000,
		},
	}

	yuriPointer := &yuri
	yuriPointer.updateName("Duy")
	yuri.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
