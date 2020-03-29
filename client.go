package hbhcl

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

const feedURL = "https://ema-hiro.hatenablog.com/feed"

// FetchFeed fetches hatena blog feed.
func FetchFeed() (*Feed, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("faield to request to hateba blog url. status: %v", resp.StatusCode)
	}

	feed := &Feed{}
	if err := xml.NewDecoder(resp.Body).Decode(feed); err != nil {
		return nil, err
	}
	return feed, nil
}
