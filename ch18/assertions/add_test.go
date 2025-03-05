package main

import "testing"

func TestAddPositiveNumbers(t *testing.T) {
    result := AddPositiveNumbersWithAssertion(5, 3)
    if result != 8 {
        t.Errorf("Expected 8, but got %d", result)
    }


    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Expected panic for negative input, but got no panic")
        }
    }()
    AddPositiveNumbersWithAssertion(-5, 3)
}

