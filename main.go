package main

import (
	"log"

	"github.com/shadyashraf174/myniceprogram/helpers"
)

const numberPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numberPool)
	intChan <- randomNumber
}

func main() {
	// pln("Hi")

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)

}

// func pln(s string) {
// 	log.Println(s)
// }
