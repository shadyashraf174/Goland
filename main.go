package main

import "fmt"

type WeaponType int

const (
	Axe WeaponType = iota // increment
	Sword
	Woodenstick
	Knife
)

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "SWORD"
	case Axe:
		return "AXE"
	}
	return ""
}

func getDomage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 70
	case Woodenstick:
		return 20
	case Knife:
		return 60
	default:
		panic("Weapon does not exist")
	}
}

func main() {
	fmt.Printf("damage of weapon (%s) (%d):\n", Axe, getDomage(Axe))
}
