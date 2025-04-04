package main

import (
	"fmt"
	"sort"
)

// Struct to hold each page's URL and its internal link count
type pageData struct {
	URL   string
	Count int
}

// Function to print a sorted report of crawled pages and their link counts
func printReport(pages map[string]int, baseURL string) {
	// Print a nicely formatted report header
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	// Convert the pages map to a slice of pageData structs for sorting
	pageList := mapToSortedSlice(pages)

	// Print each page's count and URL in the desired format
	for _, page := range pageList {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
}

// Helper function to convert and sort the pages map by count and URL
func mapToSortedSlice(pages map[string]int) []pageData {
	var pageList []pageData
	for url, count := range pages {
		pageList = append(pageList, pageData{URL: url, Count: count})
	}

	// Sort by Count in descending order; for ties, sort alphabetically by URL
	sort.Slice(pageList, func(i, j int) bool {
		if pageList[i].Count == pageList[j].Count {
			return pageList[i].URL < pageList[j].URL
		}
		return pageList[i].Count > pageList[j].Count
	})

	return pageList
}
