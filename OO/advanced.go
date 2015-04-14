package main

import (
	"log"
	"math/rand"
)

type Vehicle struct {
	Wheels                   int
	Type                     string `Vehicle:"unicycle"`
	Fuel, Consumes, Distance float32
}

type Bicycle struct {
	*Vehicle
}

type Car struct {
	*Vehicle
}

func (vehicle *Vehicle) Travel(distance float32) float32 {
	if vehicle.Fuel < 0 {
		log.Print("Out of fuel!")
	} else {
		vehicle.Distance += distance * vehicle.Consumes
		vehicle.Fuel -= vehicle.Distance
	}

	return vehicle.Fuel
}

func main() {
	car := Car{&Vehicle{4, "car", 100, 1.5, 0}}

	r := rand.New(rand.NewSource(1))
	for car.Travel(r.Float32()) > 0 {
		log.Printf("%s travelled %e, consuming %e", car.Type, car.Distance, car.Consumes)
	}

	r2 := rand.New(rand.NewSource(1))
	bike := Bicycle{&Vehicle{2, "bicycle", 100, 0.5, 0}}

	for bike.Travel(r2.Float32()) > 0 {
		log.Printf("%s travelled %e, consuming %e", bike.Type, bike.Distance, bike.Consumes)
	}
}
