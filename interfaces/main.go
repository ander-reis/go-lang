package main

import "fmt"

type vehicle interface {
	start() string
}

type car struct {
	name string
}

type motorcycle struct {
	name string
}

func (m motorcycle) start() string {
	return "the motorcycle " + m.name + " has been started"
}

func (c car) start() string {
	return "the car " + c.name + " has been started"
}

func startVehicle(v vehicle) string {
	return v.start()
}

func main() {
	c := car{"gol"}
	m := motorcycle{"harley"}

	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(m))

	//fmt.Println(c.startCar())
	//fmt.Println(m.startMotorcycle())
}
