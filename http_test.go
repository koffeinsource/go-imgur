package imgur

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
)

func testHTTPClientJSON(json string) (*http.Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-RateLimit-UserLimit", "1")
		w.Header().Set("X-RateLimit-UserRemaining", "2")
		w.Header().Set("X-RateLimit-UserReset", "3")
		w.Header().Set("X-RateLimit-ClientLimit", "4")
		w.Header().Set("X-RateLimit-ClientRemaining", "5")
		fmt.Fprintln(w, json)
	}))

	u, err := url.Parse(server.URL)
	if err != nil {
		log.Fatalln("failed to parse httptest.Server URL:", err)
	}

	http.DefaultClient.Transport = rewriteTransport{URL: u}
	return http.DefaultClient, server
}

func testHTTPClient500() (*http.Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))

	u, err := url.Parse(server.URL)
	if err != nil {
		log.Fatalln("failed to parse httptest.Server URL:", err)
	}
	http.DefaultClient.Transport = rewriteTransport{URL: u}

	return http.DefaultClient, server
}

func testHTTPClientInvalidJSON() (*http.Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")

		// some broken headers
		w.Header().Set("X-RateLimit-UserLimit", "a")
		w.Header().Set("X-RateLimit-UserRemaining", "b")
		w.Header().Set("X-RateLimit-UserReset", "c")
		w.Header().Set("X-RateLimit-ClientLimit", "d")
		w.Header().Set("X-RateLimit-ClientRemaining", "e")

		// some invalid json
		fmt.Fprintln(w, `[broken json.. :)]`)
	}))

	u, err := url.Parse(server.URL)
	if err != nil {
		log.Fatalln("failed to parse httptest.Server URL:", err)
	}
	http.DefaultClient.Transport = rewriteTransport{URL: u}

	return http.DefaultClient, server
}

type rewriteTransport struct {
	Transport http.RoundTripper
	URL       *url.URL
}

func (t rewriteTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.URL.Scheme = t.URL.Scheme
	r.URL.Host = t.URL.Host
	r.URL.Path = path.Join(t.URL.Path, r.URL.Path)
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(r)
}
