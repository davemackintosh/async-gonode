package main

import "log"

type Animal struct {
	Says string `Animal:"Nothing..."`
}

type Cat struct {
	*Animal
}

func (animal *Animal) Speak() string {
	return animal.Says
}

func main() {
	a_cat := Cat{&Animal{"Meow"}}

	log.Print(a_cat.Speak())
}
