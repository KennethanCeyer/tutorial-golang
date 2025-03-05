package main

import (
	"fmt"
	"os"
)


func main() {
    if _, err := fmt.Println("This is a standard output message"); err != nil {
        fmt.Fprintf(os.Stderr, "Error occurred: %v\n", err)
    }


    // 의도적으로 오류 발생
    if _, err := os.Open("nonexistentfile.txt"); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
    }
}
