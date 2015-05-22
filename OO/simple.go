package main

import "log"

type Component struct {
	Name string
}

type Components []Component

type Bicycle struct {
	Brand string
	Model string
	Size  float32
	Components
}

func (self Component) name() string {
	return self.Name
}

func (cats Component) Speak() string {
	return cats.Name
}

func main() {
	myComponents := Components{
		Component{"frame"},
		Component{"forks"},
		Component{"shock"},
		Component{"wheels"},
	}

	bike := Bicycle{"Scott", "YPZ2", 15.5, myComponents}

	log.Print(bike.Components[0].Speak())

	log.Print(bike)
}
