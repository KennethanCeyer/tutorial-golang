package main

import "fmt"

func main() {
	// equal
	fmt.Printf("%v\t=\ttrue == false\n", true == false) // false

	// greater than
	fmt.Printf("%v\t=\t2 > 3\n", 2 > 3) // false
	fmt.Printf("%v\t=\t\"go\" > \"c\"\n", "go" > "c") // true

	// greater than or equal
	fmt.Printf("%v\t=\t3.5 >= 3\n", 3.5 >= 3) // true
	fmt.Printf("%v\t=\t'Z' >= 'A\n", 'Z' >= 'A') // true

	// less than
	fmt.Printf("%v\t=\t-1 < -1.1\n", -1 < -1.1) // false

	// less than or equal
	fmt.Printf("%v\t=\t3.5 <= 3.5\n", 3.5 <= 3.5) // true
}
