package main

import "fmt"

const HELLO = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "world"
	}
	return HELLO + name
}

func main() {
	fmt.Println("Hello world!")
}
