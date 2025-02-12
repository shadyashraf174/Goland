package main

import "fmt"

type Entity struct {
	name    string
	id      string
	version string
	posx    int
	posy    int
}

type SpecialField struct {
	Entity
	SpecialField float64
}

type Color int

func (c Color) String() string {
	switch c {
	case CorlorBlue:
		return "BLUE"
	case ColorBlack:
		return "BLACK"
	case ColorYellow:
		return "YELLOW"
	case ColorPink:
		return "Pink"
	default:
		panic("invalid color given")
	}
}

const (
	CorlorBlue Color = iota + 1
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {

	e := Entity{
		name:    "shady",
		id:      "id 1",
		version: "version 1",
		posx:    100,
		posy:    200,
	}

	en := SpecialField{
		SpecialField: 1.23,
		Entity: Entity{
			name:    "shady",
			id:      "id 1",
			version: "version 1",
			posx:    100,
			posy:    200,
		},
	}

	fmt.Printf("this is the Entity%+v%+v\n", en.id, e)

	fmt.Println("this color is: ", ColorBlack, CorlorBlue)
}
