package main

import "log"

func main() {

	var myString string
	myString = "Green"
	log.Println("")
	log.Println("myString is set to", myString)
	changUsingPointer(&myString)
	log.Println("myString is set to", myString)
	log.Println("")
}

func changUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue

}
