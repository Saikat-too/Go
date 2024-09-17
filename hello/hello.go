package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Saikat")

	fmt.Println(message)
}
