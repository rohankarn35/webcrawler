package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {

	cfg.mu.Lock()

	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL]++
		return false
	} else {
		cfg.pages[normalizedURL] = 1
		return true
	}
}
func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("Failed to normalize URL: %v\n", err)
		return
	}

	if !cfg.addPageVisit(normalizedURL) {
		log.Printf("URL already visited: %s\n", normalizedURL)
		return // Skip already visited URLs
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("Failed to get HTML from %s: %v\n", rawCurrentURL, err)
		return
	}

	foundURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		// log.Printf("Failed to extract URLs from HTML: %v\n", err)
		return
	}

	for _, foundURL := range foundURLs {

		cfg.wg.Add(1)
		go cfg.crawlPage(foundURL)
	}
}
