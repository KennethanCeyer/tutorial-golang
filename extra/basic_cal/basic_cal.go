package main

import "fmt"

func add(a float64, b float64) float64 {
	return a + b
}

func sub(a float64, b float64) float64 {
	return a - b
}

func mul(a float64, b float64) float64 {
	return a * b
}

func div(a float64, b float64) float64 {
	return a / b
}

func main() {
	// ((2 + 3) * 4 / 3) - 6
	fmt.Printf("((2 + 3) * 4 / 3) - 6 = %.2f\n",
		sub(div(mul(add(2.0, 3.0), 4.0), 3.0), 6.0))
}
