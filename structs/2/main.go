package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	Name   string
	CanFly bool
}

func (car Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s", car.Name, car.Year, car.Color)
}

func main() {
	car1 := Car{"Corola", 2019, "Branco"}

	sCar := SuperCar{
		Car:    car1,
		CanFly: true,
		Name:   "teste",
	}
	//car2 := Car{"BMW", 2018, "Preto"}
	//fmt.Println(car1.Name, car1.Year, car1.Color)
	//fmt.Println(car2.Name, car2.Year, car2.Color)

	fmt.Println(car1.info())
	fmt.Println(sCar.Name)
}
