package main

import (
	"log"
	"net/http"
)

func main() {

	// var isTrue bool
	// isTrue = false

	// if isTrue == true {
	// 	log.Println(isTrue)
	// } else {
	// 	log.Println("nothing")
	// }

	cat := "cat"

	// if cat == "cat" {
	// 	log.Println("isTrue")
	// } else {
	// 	log.Println("isNotTrue")
	// }

	switch cat {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":
		log.Println("cat is set to dog")

	default:
		log.Println("cat is set to somthing else")
	}

}

//01:42:10
