package main

import "fmt"

const HELLO_ENGLISH = "Hello, "
const HELLO_SPANISH = "Hola, "
const HELLO_FRENCH = "Bonjour, "

func getPrefix(language string) (prefix string) {
	switch language {
	case "english":
		prefix = HELLO_ENGLISH
	case "french":
		prefix = HELLO_FRENCH
	case "spanish":
		prefix = HELLO_SPANISH
	default:
		prefix = HELLO_ENGLISH
	}
	return
}

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := getPrefix(language)
	return prefix + name
}

func main() {
	fmt.Println("Hello world!")
}
