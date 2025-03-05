package main

import "testing"

func Test_GetTimezone(t *testing.T) {
	asiaSeoul, err := GetTimezone("Asia/Seoul")
	if err != nil {
		t.Fatal(err)
	}

	asiaTokyo, err := GetTimezone("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	americaNewYork, err := GetTimezone("America/New_York")
	if err != nil {
		t.Fatal(err)
	}

	unsupportedTimezone, err := GetTimezone("Europe/London")
	if unsupportedTimezone != nil || err == nil {
		t.Fatal("Europe/London timezone name should be nil")
	}

	if asiaSeoul.Name != "KST" {
		t.Fatal("Asia/Seoul timezone name should be KST")
	}

	if asiaSeoul.UTC != 9 {
		t.Fatal("Asia/Seoul timezone UTC 9")
	}

	if asiaTokyo.Name != "JST" {
		t.Fatal("Asia/Tokyo timezone name should be JST")
	}

	if asiaTokyo.UTC != 9 {
		t.Fatal("Asia/Tokyo timezone UTC 9")
	}

	if americaNewYork.Name != "EST" {
		t.Fatal("America/New_York timezone name should be EST")
	}

	if americaNewYork.UTC != -5 {
		t.Fatal("America/New_York timezone UTC -5")
	}
}
