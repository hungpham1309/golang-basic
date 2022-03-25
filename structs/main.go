package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

/*
* in front of a type => looking for a pointer to that type
* in front of a variable => turn a pointer to value
& in front of a variable => get a pointer of a variable
*/
func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{firstName: "alex",
		lastName: "becker",
		contactInfo: contactInfo{
			email:   "hung@abc.com",
			zipCode: 3,
		},
	}

	alex.print()

	alex.updateName("hung")

	alex.print()
}
