package main

import "fmt"

func main() {
	//colors := make(map[string]string)
	colors := map[string]string{
		"white": "#ffffff",
		"black": "#000000",
	}

	colors["red"] = "#ff0000"
	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("The hex of %s is %s\n", color, hex)
	}
}
