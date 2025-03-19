package main


import "errors"


// StubDatabase는 테스트를 위한 간단한 데이터베이스 구현이다.
type StubDatabase struct {
    data map[string]string
}


// GetData는 고정된 데이터를 반환한다.
func (db *StubDatabase) GetData(key string) (string, error) {
    if value, exists := db.data[key]; exists {
        return value, nil
    }
    return "", errors.New("key not found")
}

