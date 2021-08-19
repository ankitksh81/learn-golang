package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// we can just write contackInfo without contact.
	// similar to contactInfo of type contactInfo
	contact contactInfo
}

func main() {
	// alex := person{"Alex", "Ernst"} // one way to define struct values
	// alex := person{firstName: "Alex", lastName: "Ernst"} // another way to define struct values
	// var alex person // yet another way to define struct values
	// alex.firstName = "Alex"
	// alex.lastName = "Ernst"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) // %+v list all fields and values in formated value

	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contact: contactInfo{ // contactInfo: contactInfo{....}
			email:   "jim@gmail.com",
			zipCode: 573939,
		},
	}
	// fmt.Printf("%+v", jim)
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p) // general print func with receiver func
}
