package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]User)

	me := User{
		FirstName: "shady",
		LastName:  "AShraf",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	// Slices

	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 1)

	sort.Ints(mySlice)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[6:8])

	names := []string{"shady", "Ashraf", "Ahmed"}
	log.Println(names)
	log.Println(names[0:2])
}
