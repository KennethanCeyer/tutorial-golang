package main

import "errors"


// MockAPIClient는 테스트를 위한 가짜 객체이다.
type MockAPIClient struct{}


// FetchData는 고정된 데이터를 반환한다.
func (m *MockAPIClient) FetchData(endpoint string) (string, error) {
    if endpoint == "test/valid" {
        return "Mock Data", nil
    }
    return "", errors.New("invalid endpoint")
}

