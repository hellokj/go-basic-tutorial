package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("kj")
	fmt.Println(message)
}
