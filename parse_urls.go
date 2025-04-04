package main

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseUrl string) ([]string, error) {
	baseurl, err := url.Parse(rawBaseUrl)
	if err != nil {
		return nil, err
	}
	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}
	var urls []string

	var findLinks func(*html.Node)
	findLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}
					absoluteUrl := baseurl.ResolveReference(link)
					urls = append(urls, absoluteUrl.String())
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			findLinks(child)
		}
	}
	findLinks(doc)
	if urls == nil {
		return nil, errors.New("no urls found")
	}
	return urls, nil
}
