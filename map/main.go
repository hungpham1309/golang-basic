package main

import "fmt"

func main() {

	//colors := make(map[string]string)
	//var colors = map[string]string
	colors := map[string]string{
		"red":   "#1234",
		"green": "#asdf",
		"white": "#white",
	}

	printMap(colors)

	delete(colors, "red")
}

func printMap(c map[string]string) {
	for c, h := range c {
		fmt.Println("Hex code for", c, "is", h)
	}
}
