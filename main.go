package main

import "fmt"

func main() {
	println("Task 1")

	
	
	
	var a rune
	for a = 48; a <= 57; a++ {
		fmt.Printf("%c", a)
	}

	println("\nTask 2")

	var b rune
	for b = 97; b <= 122; b++ {
		fmt.Printf("%c", b)
	}

	println("\nTask 3")

	var c rune
	for c = 57; c >= 48; c-- {
		fmt.Printf("%c", c)
	}

	println("\nTask 4")

	var d rune
	for d = 122; d >= 97; d-- {
		fmt.Printf("%c", d)
	}
}
