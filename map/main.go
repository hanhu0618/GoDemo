package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[int]string)

	// colors[10] = "#ffffff"

	// delete(colors, 10)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
