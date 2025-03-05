package main

import "fmt"

func main() {
	literalMap := map[string]string{
		"Apple":    "Fruit",
		"Orange":   "Fruit",
		"Tomato":   "Vegetable",
		"Onion":    "Vegetable",
		"Cucumber": "Cumberbatch",
	}

	for name, category := range literalMap {
		fmt.Printf("| %-10s | %-15s |\n", name, category)
	}
}
