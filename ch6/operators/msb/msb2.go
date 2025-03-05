package main

import "fmt"

func main() {
    var operand1 int32 = 29
    var operand2 int32 = ^19 + 1
    
    var result int32 = operand1 + operand2
     
    fmt.Printf("operand1: %.32b\n", operand1)
    fmt.Printf("operand2: %.32b\n", operand2)
    fmt.Printf("결과 값 (2진수): %.32b\n", result)
    fmt.Printf("결과 값은: %d", result)
}

