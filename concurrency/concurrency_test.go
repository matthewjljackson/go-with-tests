package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebsiteChecker, urls)
	}
}

func MockWebsiteChecker(url string) bool {
	return url != "fakewebsite.com"
}

func TestCheckWebsite(t *testing.T) {

	websites := []string{
		"fakewebsite.com",
		"realwebsite.com",
		"google.com",
	}

	want := map[string]bool{
		"fakewebsite.com": false,
		"realwebsite.com": true,
		"google.com":      true,
	}

	got := CheckWebsite(MockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
