package main

import (
	"fmt"
	"sync"
)

// value struct
type FetchedUrl struct {
	url map[string]bool
	mux sync.Mutex
}

func (f *FetchedUrl) Add(url string) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.url[url] = true
}

func (f FetchedUrl) Exists(url string) bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	// keyの値がない時にはboolの初期値であるfalseが返る
	return f.url[url]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fetchedUrl FetchedUrl) {
	// cacheが存在すればリクエストを送らない
	if fetchedUrl.Exists(url) || depth <= 0 {
		return
	}

	ch := make(chan []string)
	go Fetch(url, ch)
	fetchedUrl.Add(url)
	urls := <-ch

	for _, u := range urls {
		Crawl(u, depth-1, fetcher, fetchedUrl)
	}
	return
}

func main() {
	fetchedUrl := FetchedUrl{url: make(map[string]bool)}
	Crawl("http://golang.org/", 4, fetcher, fetchedUrl)
}

func Fetch(url string, ch chan []string) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		close(ch)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	ch <- urls
	close(ch)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
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
