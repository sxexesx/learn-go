package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	template = "http://books.toscrape.com/catalogue/page-50.html"
)

func main() {
	runtime.GOMAXPROCS(10)

	start := time.Now()
	httpClient := NewHttpClient()

	//
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	urls := make([]string, 0)

	// 1 ... 50
	for i := 1; i < 50; i++ {
		url, err := linkBuilder(i)
		if err != nil {
			log.Fatalln("building url")
		}
		urls = append(urls, url)
	}

	books := make([]string, 0)

	tokens := make(chan time.Time, 1)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for c := range ticker.C {
			select {
			case tokens <- c:
			case <-ctx.Done():
				log.Fatal("time exceeded")
				return
			}
		}
	}()

	for _, url := range urls {
		<-tokens

		go func() {
			titles, err := crawl(ctx, url, httpClient)
			if err != nil {
				log.Fatalln("crawling")
			}

			books = append(books, titles...)
		}()
	}

	fmt.Println(books)
	fmt.Println(len(books), " books in library")
	fmt.Println("process done for ", time.Since(start))
}

func linkBuilder(n int) (string, error) {
	basePath := "http://books.toscrape.com/catalogue"

	suffix := fmt.Sprintf("page-%d.html", n)
	return url.JoinPath(basePath, suffix)
}

type httpClient struct {
	client http.Client
}

func NewHttpClient() httpClient {
	return httpClient{
		client: http.Client{},
	}
}

func crawl(ctx context.Context, url string, client httpClient) ([]string, error) {
	println(url)
	bytes, err := client.Get(ctx, url)
	if err != nil {
		log.Fatalf("bad response from %s", url)
	}

	titles, err := findBookName(bytes)
	if err != nil {
		log.Fatalln("finding titles")
	}

	return titles, nil
}

func (h *httpClient) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	return bytes, err
}

// not realised
func findBookName(data []byte) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	titles := make([]string, 0)

	_ = doc.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		titles = append(titles, title)
	})

	return titles, nil
}
