package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {

	// Given
	urls := []string{
		"https://google.com",
		"https://twitter.com",
		"https://linkedin.com",
		"https://github.com",
		"waat://furhurterwe.geds",
	}
	var want = make(map[string]bool)
	want = map[string]bool{
		"https://google.com":      true,
		"https://twitter.com":     true,
		"https://linkedin.com":    true,
		"https://github.com":      true,
		"waat://furhurterwe.geds": false,
	}

	// When
	got := CheckWebsites(mockWebsiteChecker, urls)

	//Then
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
