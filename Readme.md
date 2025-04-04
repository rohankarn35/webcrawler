# Web Crawler in Go

A high-performance concurrent web crawler built in Go that efficiently crawls websites while respecting concurrency limits and maximum page constraints.

## Features

- **Concurrent Crawling**: Utilizes Go's goroutines for parallel processing
- **Configurable Limits**: Set maximum concurrency and page visit limits
- **Domain Scoping**: Only crawls pages within the specified domain
- **URL Normalization**: Prevents duplicate visits using URL normalization
- **Results Reporting**: Generates summary of crawled pages with link counts

## Installation

```bash
git clone https://github.com/rohankarn35/webcrawler
cd webcrawler
go mod download
```

## Usage

```bash
go run . [website_url] [max_concurrency] [max_pages]
```

### Parameters

- `website_url`: The URL of the website to crawl (e.g., https://example.com)
- `max_concurrency`: Maximum number of concurrent requests allowed
- `max_pages`: Maximum number of pages to crawl before stopping

### Example

```bash
go run . https://example.com 10 100
```

This will crawl example.com using 10 concurrent crawlers and stop after visiting 100 pages.

## Code Structure

- `main.go`: Entry point and command-line argument processing
- `models.go`: Data structures for crawler configuration
- `page_crawl.go`: Main crawling logic with concurrency control
- `normalize_url.go`: URL normalization to prevent duplicate visits
- `parse_html.go`: HTML fetching functionality
- `parse_urls.go`: URL extraction from HTML content
- `pageData.go`: Results formatting and reporting

## Testing

The project includes unit tests for core functionality:

```bash
go test -v
```

## Dependencies

- Go 1.21.6+
- golang.org/x/net package for HTML parsing

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Submit a pull request
