package main

import "fmt"

func main() {

	colors := map[string]string {
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
		"black": "#000000",
		"white": "#ffffff",
	}

	//var colors map[string]string
	//colors := make(map[string]string)
	//colors["white"] = "#ffffff"
	//colors[10] = "#000000"
	//delete(colors, "white")
	//fmt.Println(colors)

	iterate(colors)

}

func iterate(k map[string]string) {
	for k, v := range k {
		fmt.Println(k, v)
	}
}

