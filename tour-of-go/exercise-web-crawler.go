package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"util"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Crawler struct {
	CacheUrls []string
	mux       sync.Mutex
}

func (c *Crawler) Fetch(url string) (body string, urls []string, err error) {
	resp, err := http.Get(url)
	respBody, _ := ioutil.ReadAll(resp.Body)
	body = string(respBody)

	if err != nil {
		return string(body), nil, err
	}

	defer resp.Body.Close()

	receiveUrls := util.ParseUrls(body)

	for _, val := range receiveUrls {
		urls = append(urls, val[1])
	}

	return body, urls, nil
}

func (c *Crawler) AddCacheUrls(u string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.CacheUrls = append(c.CacheUrls, u)
}

func (c *Crawler) GetCacheUrls() []string {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.CacheUrls
}

func Crawl(url string, depth int, fetcher Fetcher) {

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: url => %s, title => %q\n", url, util.ParseTitle(body))
	for _, u := range urls {
		if !strings.HasPrefix(u, "http") {
			rootUrl, _ := util.GetRootUrl(url)
			u = rootUrl + u
		}

		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	// https://www.aptpod.co.jp/index.html
	accessUrl := "https://www.aptpod.co.jp/index.html"
	c := new(Crawler)
	Crawl(accessUrl, 1, c)
}
