package main

import (
	"fmt"
)

func pad(l int, c string) string {
	l = 4 - l
	s := ""
	for i := 0; i < l; i++ {
		s += c
	}
	return s
}

// Crawl使用fetcher从某个URL开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool) {
	var c func(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool)
	c = func(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool) {
		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			fmt.Println(err)
			return
		}

		history[url] = true

		ch <- fmt.Sprintf("%sfound %s %q\n **** %v\n", pad(depth, "-"), url, body, urls)
		for _, u := range urls {
			if !history[u] {
				history[u] = true
				c(u, depth-1, fetcher, ch, history)
			} else {
				fmt.Printf(">>>%s is repetitive\n", u)
			}
		}
		return
	}

	c(url, depth, fetcher, ch, history)

	close(ch)
}

func main() {
	ch := make(chan string)
	history := make(map[string]bool)
	go Crawl("http://golang.org/", 4, fetcher, ch, history)
	for s := range ch {
		fmt.Println(s)
	}

	fmt.Println("============= history ================")
	fmt.Println(history)
}

type Fetcher interface {
	// Fetch 返回URL的body内容，并且将在这个页面上找到的URL放到一个slice中。
	Fetch(url string) (body string, urls []string, err error)
}

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher是返回若干结果的Fetcher.
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found %s", url)
}

// fetcher是填充后的fakeFetcher
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

/*
found http://golang.org/ "The Go Programming Language"
 **** [http://golang.org/pkg/ http://golang.org/cmd/]

-found http://golang.org/pkg/ "Packages"
 **** [http://golang.org/ http://golang.org/cmd/ http://golang.org/pkg/fmt/ http://golang.org/pkg/os/]

>>>http://golang.org/ is repetitive
not found http://golang.org/cmd/
>>>http://golang.org/ is repetitive
>>>http://golang.org/pkg/ is repetitive
--found http://golang.org/pkg/fmt/ "Package fmt"
 **** [http://golang.org/ http://golang.org/pkg/]

--found http://golang.org/pkg/os/ "Package os"
 **** [http://golang.org/ http://golang.org/pkg/]

>>>http://golang.org/ is repetitive
>>>http://golang.org/pkg/ is repetitive
>>>http://golang.org/cmd/ is repetitive
============= history ================
map[http://golang.org/cmd/:true http://golang.org/pkg/fmt/:true http://golang.org/pkg/os/:true http://golang.org/:true http://golang.org/pkg/:true]
*/

// https://github.com/davidhoo/go-tour/blob/master/exercise-web-crawler.go
