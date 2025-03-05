package main


import "testing"


func TestStubDatabase(t *testing.T) {
    // Stub 데이터베이스 초기화
    stubDB := &StubDatabase{
        data: map[string]string{
            "key1": "value1",
            "key2": "value2",
        },
    }


    // "key1"에 대한 결과 검증
    result, err := stubDB.GetData("key1")
    if err != nil || result != "value1" {
        t.Errorf("Expected 'value1', got '%s', error: %v", result, err)
    }


    // 존재하지 않는 키에 대한 결과 검증
    result, err = stubDB.GetData("key3")
    if err == nil || result != "" {
        t.Errorf("Expected error for 'key3', got '%s', error: %v", result, err)
    }
}

