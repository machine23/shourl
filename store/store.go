package store

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"encoding/base64"
)

var storeIndex int

type engine interface {
	getID(value string) string
	get(id string) string
	put(id, value string) error
}

// Store is a store for keeping URLs.
type Store struct {
	storeType string
	engine
}

var timestamp = func() int64 {
	return time.Now().Unix()
}

// New creates store.
// Possible types of store: "memory"
func New(storeType string) (*Store, error) {
	switch storeType {
	case "memory":
		return &Store{storeType: storeType, engine: makeMemoryStore()}, nil
	default:
		return nil, fmt.Errorf("unsupported store type: %s", storeType)
	}
}

func newID(v string) string {
	storeIndex++
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(strconv.Itoa(storeIndex)))
}

// Type returns a type of the store
func (s Store) Type() string {
	return s.storeType
}

// AddURL puts a given url to the store and returns its id.
func (s Store) AddURL(url string) (string, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return "", fmt.Errorf("Empty url is not allowed")
	}
	id := s.getID(url)
	if id != "" {
		return id, nil
	}

	id = newID(url)
	if err := s.put(id, url); err != nil {
		return "", err
	}
	return id, nil
}

// GetURL returns url by the given id. If there is no url with this id,
// it returns an empty string.
func (s Store) GetURL(id string) string {
	return s.get(id)
}
