package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	ch := make(chan UrlAndTitle)
	go crawlRecursive(url, depth, fetcher, ch)
	for urlAndTitle := range ch {
		fmt.Printf("found: %s %q\n", urlAndTitle.url, urlAndTitle.title)
	}
}

func crawlRecursive(url string, depth int, fetcher Fetcher, ch chan UrlAndTitle) {
	defer close(ch)
	if depth <= 0 {
		return
	}

	if alreadyFetchedUrls.isSet(url) {
		return
	}

	alreadyFetchedUrls.Add(url)

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)

		return
	}

	ch <- UrlAndTitle{
		url,
		body,
	}

	subChannels := make([]chan UrlAndTitle, len(urls))
	for i, u := range urls {
		subChannels[i] = make(chan UrlAndTitle)
		go crawlRecursive(u, depth-1, fetcher, subChannels[i])
	}

	for i := range subChannels {
		for j := range subChannels[i] {
			ch <- j
		}
	}

	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

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

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type AlreadyFetchedUrls struct {
	mu   sync.Mutex
	urls map[string]int
}

var alreadyFetchedUrls = AlreadyFetchedUrls{
	urls: make(map[string]int),
}

func (alreadyFetchedUrls *AlreadyFetchedUrls) Add(key string) {
	alreadyFetchedUrls.mu.Lock()
	alreadyFetchedUrls.urls[key] = 0
	alreadyFetchedUrls.mu.Unlock()
}

func (AlreadyFetchedUrls *AlreadyFetchedUrls) isSet(key string) bool {
	alreadyFetchedUrls.mu.Lock()
	defer alreadyFetchedUrls.mu.Unlock()
	_, exists := alreadyFetchedUrls.urls[key]

	return exists
}

type UrlAndTitle struct {
	url string
	title string
}