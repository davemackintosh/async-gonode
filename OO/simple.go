package main

import "log"

// Animal njwigniwg wrbighwrog wrgh wgw
type Animal struct {
	Says string `Animal:"Nothing..."`
}

// I'm a Cat
type Cat struct {
	*Animal
}

func (animal *Animal) Speak() string {
	return animal.Says
}

func main() {
	a_cat := Cat{&Animal{"Meow"}}
	animal := new(Animal)

	log.Print(a_cat)
	log.Print(animal)

	log.Print(a_cat.Speak())
}
