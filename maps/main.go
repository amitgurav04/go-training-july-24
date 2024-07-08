package main

import "fmt"

func main() {

	//var colors map[string]string

	colors := make(map[string]string)

	/*colors := make(map[string]string) {

	}
	colors = {
		"red":    "#ff0000",
		"green":  "#00ff00",
		"blue":   "#0000ff",
		"orange": "#ff0000",
	}*/

	colors["red"] = "#FF1111"
	colors["blue"] = "#00FF00"
	colors["cyan"] = "#0000FF"
	colors["white"] = "#FFFFFF"
	fmt.Println("=================Before================")
	printMap(colors)
	updateMap(colors)
	fmt.Println("=================After================")
	printMap(colors)

}

func updateMap(colors map[string]string) {
	colors["red"] = "#FF0000"
	colors["orange"] = "#00FF00"
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code of ", key, "is", value)
	}
}
