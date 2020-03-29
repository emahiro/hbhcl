package hbhcl

import (
	"net/url"
	"time"
)

// Feed represents hatena blog feed in Atom format.
type Feed struct {
	Title   string  `xml:"title"`
	Link    string  `xml:"link"`
	Author  string  `xml:"author"`
	Entries []Entry `xml:"entry"`
}

// Entry represents hatena blog entry.
type Entry struct {
	Title     string    `xml:"title"`
	Link      *url.URL  `xml:"link"`
	ID        string    `xml:"id"`
	Published time.Time `xml:"published"`
	Updated   time.Time `xml:"updated"`
	Summary   string    `xml:"summary"`
	Content   string    `xml:"content"`
	Author    string    `xml:"author"`
	Category  string    `xml:"category"`
}
