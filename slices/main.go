package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	//slice = append(slice, 10, 3, 5, 1)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//for i := 0; i < 20; i++ {
	//	slice = append(slice, i)
	//	fmt.Println("Len: ", len(slice), "Cap: ", cap(slice))
	//}

	//sliceTest := slice
	//slice = append(slice, 1, 2, 3, 4, 5)
	//slice[0] = 10
	//fmt.Println(slice)
	//fmt.Println(sliceTest)

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
	}
	fmt.Println(sliceString)
	fmt.Println(sliceString[0])   // hello
	fmt.Println(sliceString[:2])  // hello world
	fmt.Println(sliceString[1:2]) // world
	fmt.Println(sliceString[2:4]) // much better
	fmt.Println(sliceString[2:])  // muche better ....
}
