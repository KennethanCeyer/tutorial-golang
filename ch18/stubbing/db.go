package main


import "errors"


type Database interface {
    GetData(key string) (string, error)
}


type RealDatabase struct {
    data map[string]string
}


func (db *RealDatabase) GetData(key string) (string, error) {
    if value, exists := db.data[key]; exists {
        return value, nil
    }
    return "", errors.New("key not found")
}

