package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Omar")
	fmt.Println(message)
}
