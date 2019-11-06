package main

// 'Droid Sans Mono', 'monospace', monospace, 'Droid Sans Fallback'

import "fmt"

func main() {

	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "ffffff",
	}

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	for key, value := range colors {
		fmt.Println("chave= " + key + " - value= " + value)
	}
}
