package shortener

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/machine23/shourl/store"
)

// type Shortener interface {
// 	Shorten(url string) string
// 	Resolve(url string) string
// }

// Shortener implements logic to shorten URLs.
type Shortener struct {
	db     *store.Store
	domain string
}

// New returns a new object of Shortener.
func New(domain string) (*Shortener, error) {
	db, err := store.New("memory")
	if err != nil {
		return nil, err
	}
	domain = strings.TrimSpace(domain)
	if domain == "" {
		return nil, fmt.Errorf("Empty domain is not allowed")
	}
	if !strings.HasPrefix(domain, "http") {
		domain = "http://" + domain
	}
	shortener := &Shortener{db, domain}
	return shortener, nil
}

// Shorten gets a full url, saves it to some store and returns a shortened url.
func (s Shortener) Shorten(uri string) string {
	// Is this check needed here?
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return ""
	}

	id, err := s.db.AddURL(uri)
	if err != nil {
		return ""
	}

	return s.domain + "/" + id
}

// Resolve gets a shortened url and returns a full url or an empty string
// if the shortened url is not found in the store.
func (s Shortener) Resolve(uri string) string {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return ""
	}
	id := strings.TrimLeft(u.Path, "/")
	return s.db.GetURL(id)
}
