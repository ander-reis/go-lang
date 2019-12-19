package main

import (
	"fmt"
	"packages/car"
)

func main() {
	car := car.Car{"Gol", "black"}

	fmt.Println(car.Start())
}
