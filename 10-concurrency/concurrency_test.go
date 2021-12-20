package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	time.Sleep(time.Millisecond * 20)

	if url == "http://wrongurl" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	t.Run("check websites", func(t *testing.T) {
		urls := []string{
			"http://somesite.localhost",
			"http://anothersite.localhost",
			"http://wrongurl",
		}

		want := map[string]bool{
			"http://somesite.localhost":    true,
			"http://anothersite.localhost": true,
			"http://wrongurl":              false,
		}

		got := CheckWebsites(mockWebsiteChecker, urls)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v , but got %v", want, got)
		}
	})
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "http://some.url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(mockWebsiteChecker, urls)
	}
}
