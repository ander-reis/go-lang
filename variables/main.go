package main

import (
	"fmt"
)

const (
	aa string = "oi"
	bb        = 66
	cc        = 77
)

const xvz int = 1333

func main() {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	f := `Uooo`

	const xpto = 10

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
}
