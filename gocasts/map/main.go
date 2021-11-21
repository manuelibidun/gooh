package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"black": "#000000",
	}
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	delete(colors, "green")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, "=>", hex)
	}
}
