package concurrency

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds" 
}

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

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !cmp.Equal(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
