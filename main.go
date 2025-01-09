package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "shady"

	myVar2 := myStruct{
		FirstName: "Ashraf",
	}

	log.Println(myVar.printFirstName())
	log.Println(myVar2.printFirstName())

}
