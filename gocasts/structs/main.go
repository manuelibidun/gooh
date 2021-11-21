package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	// simple init
	//me := person{"Alex", "Anderson"}

	//field:value init
	me := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}
	fmt.Println(me)

	// delayed init and formatted printing
	var alex person
	fmt.Printf("%+v\n", alex)
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alexo@mail.com"
	fmt.Printf("%+v\n", alex)

	// embedded struct and receiver func
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@mail.com",
			zipCode: 23401,
		},
	}
	jim.print()

	//jimPointer := &jim
	//println(jimPointer)
	//fmt.Println(*jimPointer)
	jim.updateName("Jimmy")
	jim.print()
}
