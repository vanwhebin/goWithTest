package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowwStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	// 普通串行调用
	// for i := 0; i < b.N; i++ {
	// 	CheckWebsites(slowwStubWebsiteChecker, urls)
	// }

	// 协程并发调用
	for i := 0; i < b.N; i++ {
		CheckWebsitesWithConcurrencyChannel(slowwStubWebsiteChecker, urls)
	}
}

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

	// got := CheckWebsites(mockWebsiteChecker, websites)
	// got := CheckWebsitesWithConcurrency(mockWebsiteChecker, websites)
	got := CheckWebsitesWithConcurrencyChannel(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
