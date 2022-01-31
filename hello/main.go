package main

import "fmt"

const HELLO = "Hello, "

func hello(name string) string {
	return HELLO + name
}

func main() {
	fmt.Println("Hello world!")
}
