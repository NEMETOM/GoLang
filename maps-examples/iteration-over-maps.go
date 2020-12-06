//
// This code iterates over the defined map and prints HEX code for each color defined in the map
//
package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":"#ff0000",
		"green":"#4bf745",
		"white":"#ffffff",
	}
	
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
