package main

import "testing"

func TestFetchData_Valid(t *testing.T) {
    client := &MockAPIClient{}

    result, err := client.FetchData("test/valid")
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if result != "Mock Data" {
        t.Errorf("Expected 'Mock Data', got '%s'", result)
    }
}

func TestFetchData_Invalid(t *testing.T) {
    client := &MockAPIClient{}

    _, err := client.FetchData("test/invalid")
    if err == nil {
        t.Error("Expected error for invalid endpoint, got nil")
    }
}

