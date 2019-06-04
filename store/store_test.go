package store

import (
	"reflect"
	"testing"
)

func init() {
	timestamp = func() int64 {
		return 1257894000
	}
}

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

func Test_newID(t *testing.T) {
	tests := []struct {
		val  string
		want string
	}{
		{"asdf", "4af9f0705129f3bd"},
		{"https://google.com", "4af9f0705b1a2675"},
	}
	for _, tt := range tests {
		t.Run(tt.val, func(t *testing.T) {
			if got := newID(tt.val); got != tt.want {
				t.Errorf("newID('%v') = '%v', want '%v'", tt.val, got, tt.want)
			}
		})
	}
}

func TestStore_AddURL(t *testing.T) {
	store, err := New("memory")
	if err != nil {
		t.Fatal("Failed to make a memory store")
	}
	url := "https://google.com"
	expectedID := "4af9f0705b1a2675"
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
