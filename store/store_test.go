package store

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	// test memory
	want := &Store{storeType: "memory", engine: &memory{}}
	got, err := New("memory")
	if err != nil {
		t.Fatalf("Failed to make memory store engine: %v", err)
	}
	if got.engine == nil {
		t.Fatal("New() returns nil for the store field")
	}
	storeType := reflect.TypeOf(got.engine)
	if storeType != reflect.TypeOf(want.engine) {
		t.Fatalf(
			"New(\"memory\") returns an unexpected type of store: %v",
			storeType,
		)
	}
	// test unexpected
	_, err = New("unknown")
	if err == nil {
		t.Fatalf("Expected to get error, but didn't")
	}
}

func TestStore_AddURL(t *testing.T) {
	store, err := New("memory")
	if err != nil {
		t.Fatal("Failed to make a memory store")
	}
	url := "https://google.com"
	expectedID := "MQ"
	id, err := store.AddURL(url)
	if err != nil {
		t.Fatalf("AddURL error: %v", err)
	}
	if id != expectedID {
		t.Fatalf("AddURL returned '%v', expect '%v'", id, expectedID)
	}
	if store.get(id) != url {
		t.Fatalf("AddURL didn't put the url into the store")
	}

	// Should return the same ID for the same url
	secondID, err := store.AddURL(url)
	if err != nil {
		t.Fatalf("AddURL error adding the same url: %v", err)
	}
	if secondID != id {
		t.Fatalf("Should return same ID for the same url: got %v, want %v", secondID, id)
	}

	_, err = store.AddURL("")
	if err == nil {
		t.Fatalf("AddURL for an empty url should return an error, but doesn't")
	}

}

func TestStore_GetURL(t *testing.T) {
	store, err := New("memory")
	if err != nil {
		t.Fatal("Failed to make a memory store")
	}
	url := "https://yandex.ru"
	id, err := store.AddURL(url)
	if err != nil {
		t.Fatal("Failed to add a url to the store")
	}
	res := store.GetURL(id)
	if res != url {
		t.Fatalf("GetURL() returned '%v', expect '%v'", res, url)
	}
}

func TestStore_Type(t *testing.T) {
	st, err := New("memory")
	if err != nil {
		t.Fatal("Failed to make a memory store")
	}
	res := st.Type()
	if res != "memory" {
		t.Fatalf("Failed to return a type of the store. Got: '%v', want: 'memory'", res)
	}
}
