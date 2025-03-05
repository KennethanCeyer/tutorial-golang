package main


import "testing"


func TestMul(t *testing.T) {
    result := Mul(2, 3)
    if result != 6 {
        t.Errorf("Expected 6, but got %d", result)
    }


    result = Mul(5, 0)
    if result != 0 {
        t.Errorf("Expected 0, but got %d", result)
    }
}

