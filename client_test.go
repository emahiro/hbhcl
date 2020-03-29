package hbhcl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"sync"
	"testing"
)

type mockTransport struct {
	mu        sync.Mutex
	URL       string
	Transport http.RoundTripper
}

func (t *mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	u, err := url.Parse(t.URL)
	if err != nil {
		return nil, err
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	r.URL = u
	if t.Transport == nil {
		t.Transport = http.DefaultTransport
	}
	return t.Transport.RoundTrip(r)
}

func TestFetchFeed(t *testing.T) {
	g, err := ioutil.ReadFile(filepath.Join("testdata", t.Name()+".golden"))
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		resp    []byte
		wantErr error
	}{
		{
			name:    "success to get Hatena Blog Feed",
			resp:    g,
			wantErr: nil,
		},
		{
			name:    "failed to parse Hatena Blog Feed",
			resp:    []byte(`<feed `),
			wantErr: fmt.Errorf("hoge"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(tt.resp)
			}))
			defer ts.Close()

			http.DefaultTransport = &mockTransport{
				Transport: http.DefaultTransport,
				URL:       ts.URL,
			}

			c := NewClient("test")
			if _, err := c.FetchFeed(); tt.wantErr == nil && err != nil {
				t.Fatalf("wantErr is %v, but actual is %v", tt.wantErr, err)
			}
		})
	}
}
