package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u := uuid.Must(uuid.NewV4())
	fmt.Printf("hello, %s\n", u)
}
