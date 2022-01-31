package main

import "fmt"

func repeat(char string) string {
	var iterations string
	for i := 0; i < 5; i++ {
		iterations = iterations + char
	}
	return iterations
}

func main() {
	fmt.Println(repeat("a"))
}
