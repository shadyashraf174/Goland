package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	// for i := 0; i <= 10; i++ {

	// }

	// mySlice := []string{"shady", "ahmed", "ashraf"}
	// for _, x := range mySlice {
	// 	log.Println(x)
	// }

	// myMap := make(map[string]string)
	// myMap["dog"] = "dog"
	// myMap["cat"] = "cat"
	// myMap["fish"] = "fish"

	// for i, x := range myMap {
	// 	log.Println(i, x)
	// }

	var mySlice []User
	u1 := User{
		FirstName: "shady",
		LastName:  "Ashraf",
	}
	u2 := User{
		FirstName: "Ahmed",
		LastName:  "Ashraf",
	}

	mySlice = append(mySlice, u1)
	mySlice = append(mySlice, u2)

	for _, x := range mySlice {
		log.Println(x.FirstName)
	}

}
