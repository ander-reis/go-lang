package main

import "fmt"

func funcName(a int) int {
	return a * a
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func moreReturns(a string, b string) (string, string) {
	return a, b
}

func variadicFunc(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}
	return res
}

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {
	fmt.Println(namedReturn("anderson"))
	x, y := moreReturns("Anderson", "Reis")
	fmt.Println(x, y)
	fmt.Println(variadicFunc(1, 2, 5, 10))

	z := 0
	add := func() int {
		z += 2
		return z
	}
	fmt.Println(add())
	fmt.Println(add())
}
