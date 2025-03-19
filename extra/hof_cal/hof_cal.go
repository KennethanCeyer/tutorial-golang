package main

import "fmt"

func cal(
	c func(a float64, b float64) float64,
	a float64,
	b float64) interface{} {
	return func(args ...interface{}) interface{} {
		if len(args) == 0 {
			return c(a, b)
		}

		if len(args) > 2 {
			panic("arguments should less than 3")
		}

		ac := args[0].(func(float64, float64) float64)
		f := args[1].(float64)

		return cal(ac, c(a, b), f)
	}
}

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
	addResult := cal(add, 2.0, 3.0).(func(
		args ...interface{}) interface{})
	mulResult := addResult(mul, 4.0).(func(
		args ...interface{}) interface{})
	divResult := mulResult(div, 3.0).(func(
		args ...interface{}) interface{})
	subResult := divResult(sub, 6.0).(func(
		args ...interface{}) interface{})
	result := subResult()
	fmt.Printf("((2 + 3) * 4 / 3) - 6 = %.2f\n", result)
}
