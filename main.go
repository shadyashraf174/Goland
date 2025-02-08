package main

import "fmt"

const (
	version = 1
	keylen  = 10
)

type Player struct {
	name        string
	health      int
	attackPower float64
}

var x = 10

var (
	floatVar32 float32 = 1.0
	floatVar64 float64 = 0.1
	name               = "foo"
	firstName  string  = "foo"
	lastName   string
)

func (player Player) getHealth() int {
	return player.health
}

type Weapon string

func getweapon(weapon Weapon) string {
	return string(weapon)
}

func main() {

	numbers := []int{1, 2, 3, 4}
	// users := map[string]int{}
	// users := map[string]int{
	// 	"shady": 34,
	// }

	numbers = append(numbers, 5)
	users := make(map[string]int)

	users["foo"] = 10
	users["bar"] = 30

	name := 19
	fmt.Println(x)

	fmt.Println(name)
	fmt.Println(version)

	// Player := Player{
	// 	name:        "captin jack",
	// 	health:      100,
	// 	attackPower: 45.1,
	// }

	age, ok := users["baz"]
	if !ok {
		fmt.Println("baz not exist")
	} else {
		fmt.Println("exist in the map: ", age)
	}

	fmt.Println(numbers)
}
