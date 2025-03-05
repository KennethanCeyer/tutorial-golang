package main

import (
	"testing"
	"time"
)

func Test_ParseDate(t *testing.T) {
	value1 := "2020-06-01"
	value2 := "2021-05-01"
	value3 := "2030-01-01"

	errorValue1 := "2015/01/02"
	errorValue2 := "2020-13-01"

	zeroValue := time.Time{}

	result1, err := ParseDate(value1)
	if err != nil {
		t.Fatal(err)
	}

	result2, err := ParseDate(value2)
	if err != nil {
		t.Fatal(err)
	}

	result3, err := ParseDate(value3)
	if err != nil {
		t.Fatal(err)
	}

	errorResult1, err := ParseDate(errorValue1)
	if err == nil {
		t.Fatalf("should return the error with '%s' value",
			errorValue1)
	}

	errorResult2, err := ParseDate(errorValue2)
	if err == nil {
		t.Fatalf("should return the error with '%s' value",
			errorValue1)
	}

	expected1 := time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)
	expected2 := time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC)
	expected3 := time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)

	if result1 != expected1 {
		t.Fatalf("%v should be same as %v", result1, expected1)
	}

	if result2 != expected2 {
		t.Fatalf("%v should be same as %v", result2, expected2)
	}

	if result1 != expected1 {
		t.Fatalf("%v should be same as %v", result3, expected3)
	}

	if errorResult1 != zeroValue {
		t.Fatalf("%v should be nil", errorResult1)
	}

	if errorResult2 != zeroValue {
		t.Fatalf("%v should be nil", errorResult2)
	}
}
