package main

import "fmt"

func showLogicalOperator() {
	// AND
	fmt.Print("AND operators ====\n")
	fmt.Printf("%t && %t = %t\n", true, true, true && true)
	fmt.Printf("%t && %t = %t\n", true, false, true && false)
	fmt.Printf("%t && %t = %t\n", false, true, false && true)
	fmt.Printf("%t && %t = %t\n", false, false, false && false)
	fmt.Print("====================\n\n")

	// OR
	fmt.Print("OR operators ====\n")
	fmt.Printf("%t || %t = %t\n", true, true, true || true)
	fmt.Printf("%t || %t = %t\n", true, false, true || false)
	fmt.Printf("%t || %t = %t\n", false, true, false || true)
	fmt.Printf("%t || %t = %t\n", false, false, false || false)
	fmt.Print("====================\n\n")

	// NOT
	fmt.Print("NOT operators ====\n")
	fmt.Printf("!%t = %t\n", true, !true)
	fmt.Printf("!%t = %t\n", false, !false)
	fmt.Print("====================\n\n")
}

func main() {
	showLogicalOperator()
}
