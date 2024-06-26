package main

import (
	"reflect"
	"testing"
	"time"
)

// helper
func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowWebCheker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// test
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWeb(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

// bench mark
func BenchmarkCheckWeb(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a urls"
	}
	for i := 0; i < b.N; i++ {
		CheckWeb(slowWebCheker, urls)
	}
}
