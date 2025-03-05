package main


func Mul(a, b int) int {
    result := 0
    for i := 0; i < b; i++ {
        result = Add(result, a)
    }
    return result
}

