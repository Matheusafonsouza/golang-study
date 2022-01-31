package main

import "fmt"

func repeat(char string, times int) string {
	var iterations string
	for i := 0; i < times; i++ {
		iterations += char
	}
	return iterations
}

func main() {
	fmt.Println(repeat("a", 6))
}
