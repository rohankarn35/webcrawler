package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawUrl string) (string, error) {
	resp, err := http.Get(rawUrl)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch %s", rawUrl)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil

}
