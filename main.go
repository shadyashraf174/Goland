package main

import (
	"log"

	"github.com/shadyashraf174/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myvar helpers.SomeType

	myvar.TypeName = "Shady"
	log.Println(myvar.TypeName)
}

// 02:04:35
