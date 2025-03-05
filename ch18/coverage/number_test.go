package main

import "testing"

func TestCheckEvenOdd(t *testing.T) {
    if CheckEvenOdd(2) != "Even" {
        t.Errorf("Expected 'Even', got something else")
    }
}

