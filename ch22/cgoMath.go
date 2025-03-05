package main


// #include <math.h>
import "C"
import "fmt"


func main() {
    result := C.sqrt(16)
    fmt.Printf("C 라이브러리의 sqrt(16): %f\n", result)
}

