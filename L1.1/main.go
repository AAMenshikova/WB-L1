package main

import (
	"fmt"
)

type Human struct {
	name    string
	surname string
	age     int
	height  int
	gender  string
}

func (h *Human) grow() {
	h.height++
}

func (h *Human) birthDay() {
	h.age++
}

func (h Human) say(phrase string) {
	fmt.Println("human", phrase)
}

type Action struct {
	Human
	typeOfAction string
}

func (a Action) say(phrase string) {
	fmt.Println("action", phrase)
}

func main() {
	a := Action {
		typeOfAction: "walk",
		Human: Human {
			name:    "Egor",
			surname: "Krid",
			age:     31,
			height:  185,
			gender:  "male",
		},
	}
	fmt.Println(a)
	a.say("Hello!")
	a.Human.say("Hello!")
	a.grow()
	a.birthDay()
	fmt.Println(a)
}
