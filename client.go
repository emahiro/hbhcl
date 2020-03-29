package hbhcl

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

const feedURL = "https://%s.hatenablog.com/feed"

// Client is hatena blog http Client per user.
type Client struct {
	hcl    *http.Client
	UserID string
}

// NewClient creates Client instance.
func NewClient(userID string) *Client {
	return &Client{
		hcl:    http.DefaultClient,
		UserID: userID,
	}
}

// FetchFeed fetches hatena blog feed.
func (c *Client) FetchFeed() (*Feed, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(feedURL, c.UserID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.hcl.Do(req)
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
