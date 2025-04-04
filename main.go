package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	log.Print("started server")
	argsWithProgram := os.Args
	if len(argsWithProgram) == 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	website := argsWithProgram[1]
	maxConcurrency := argsWithProgram[2]
	maxPages := argsWithProgram[3]

	maxConcurrencyInt, err := strconv.Atoi(maxConcurrency)
	if err != nil {
		fmt.Println("Invalid max concurrency value")
		os.Exit(1)
	}

	maxPagesInt, err := strconv.Atoi(maxPages)
	if err != nil {
		fmt.Println("Invalid max pages value")
		os.Exit(1)
	}

	// fmt.Println("HTML Content:", htmlContent)
	// // pages := make(map[string]int)

	// // // Start the crawl with the base URL
	// // crawlPage("https://gobyexample.com/", "https://gobyexample.com/command-line-arguments", pages)

	// // // Print out the pages map
	// // for page, count := range pages {
	// // 	fmt.Printf("Page: %s, Visit Count: %d\n", page, count)
	// // }
	// cfg := &config{
	// 	pages:              make(map[string]int),
	// 	concurrencyControl: make(chan struct{}, maxConcurrencyInt),
	// 	wg:                 &sync.WaitGroup{},
	// 	mu:                 &sync.Mutex{},
	// 	maxPages:           maxPagesInt,
	// }

	// Set the base URL for the website we want to crawl
	baseURL := website

	// Parse the base URL
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Failed to parse base URL: %v\n", err)
		return
	}

	// Set the baseURL in the config

	cfg := &config{
		pages:              make(map[string]int),
		concurrencyControl: make(chan struct{}, maxConcurrencyInt),
		wg:                 &sync.WaitGroup{},
		mu:                 &sync.Mutex{},
		maxPages:           maxPagesInt,
		baseURL:            parsedBaseURL,
	}
	// Add the initial URL to start crawling

	fmt.Printf("starting crawl of: %s...\n", baseURL)
	// Start the crawl (this will call crawlPage in a goroutine)
	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	// Print the pages that were crawled
	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
	// printReport(cfg.pages, baseURL)

	os.Exit(0)

}
