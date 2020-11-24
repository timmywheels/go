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

func main() {
	tim := person{
		firstName: "Tim",
		lastName: "Wheeler",
		contact:contactInfo{
			email:   "tim@wheeler.com",
			zipCode: 90210,
		},
	}
	//andrea := person{
	//	firstName: "Andrea",
	//	lastName:  "Wheeler",
	//	contact: contactInfo{
	//		email:   "andrea@wheeler.com",
	//		zipCode: 90210,
	//	},
	//}

	//timPointer := &tim
	//timPointer.print()

	//var andrea person
	//fmt.Println(andrea)
	//andrea.print()
	//tim.print()
	tim.updateName("Timothy")
	tim.print()

	//fmt.Println(andrea)

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}