package main

import "fmt"

func main() {

	users := map[string]int{
		"foo": 1,
		"bar": 2,
		"Ali": 3,
		"Bob": 4,
		"brr": 5,
	}

	//names := []string{"a", "b", "c", "d", "f"}

	for key, value := range users {
		fmt.Printf("key %s value %d\n", key, value)
	}
}
