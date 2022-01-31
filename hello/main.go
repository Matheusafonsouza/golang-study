package main

import "fmt"

const HELLO_ENGLISH = "Hello, "
const HELLO_SPANISH = "Hola, "

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "spanish" {
		return HELLO_SPANISH + name
	}
	return HELLO_ENGLISH + name
}

func main() {
	fmt.Println("Hello world!")
}
