package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var fm = log.New(os.Stdout, "", 0)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	m := map[string]bool{url: true}
	var mx sync.Mutex
	var wg sync.WaitGroup
	var c2 func(string, int)

	c2 = func(url string, depth int) {
		defer wg.Done()
		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fm.Printf("Found: %s %q\n", url, body)
		mx.Lock()
		for _, u := range urls {
			if !m[u] {
				m[u] = true
				wg.Add(1)
				go c2(u, depth-1)
			}
		}
		mx.Unlock()
	}
	wg.Add(1)
	c2(url, depth)
	wg.Wait()
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakefetcher is the fetcher that returns canned results

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)

}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
