package main


import "testing"


func TestProcessInput(t *testing.T) {
    result := ProcessInput("")
    if result != -1 {
        t.Errorf("Expected -1 for empty input, but got %d", result)
    }


    result = ProcessInput("hello")
    if result != 5 {
        t.Errorf("Expected 5 for input 'hello', but got %d", result)
    }
}

