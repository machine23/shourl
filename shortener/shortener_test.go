package shortener

import (
	"reflect"
	"testing"

	"github.com/machine23/shourl/store"
)

func TestNew(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		args    args
		want    *Shortener
		wantErr bool
	}{
		{
			name:    "Empty domain is not allowed",
			args:    args{""},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check a protocol name 1",
			args: args{"otus.ru"},
			want: &Shortener{
				db:     &store.Store{},
				domain: "http://otus.ru",
			},
			wantErr: false,
		},
		{
			name: "Check a protocol name 2",
			args: args{"http://otus.ru"},
			want: &Shortener{
				domain: "http://otus.ru",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.domain)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("New() unexpected error = %v", err)
				}
				return
			}
			if err == nil && tt.wantErr {
				t.Errorf("No error returned when expected")
				return
			}
			if got == nil {
				t.Fatal("New() unexpectedly returns nil for shortener")
			}
			if reflect.TypeOf(got.db) != reflect.TypeOf(tt.want.db) {
				t.Fatalf(
					"New() type of db: %v, want: %v",
					reflect.TypeOf(got.db),
					reflect.TypeOf(tt.want.db),
				)
			}
			if got.domain != tt.want.domain {
				t.Fatal("New() didn't add http to domain")
			}
		})
	}
}

func TestShortener_Shorten(t *testing.T) {
	tests := []struct {
		name  string
		full  string
		short string
	}{
		{"Full url is empty", "", ""},
		{"Full url is not a valid url", "<1afaa>", ""},
	}

	sh, err := New("otus.ru")
	if err != nil {
		t.Fatal("Failed to create a new shortener")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sh.Shorten(tt.full); got != tt.short {
				t.Errorf("Shortener.Shorten() = %v, want %v", got, tt.short)
			}
		})
	}
}

func TestShortener_Resolve(t *testing.T) {
	sh, _ := New("otus.ru")
	full := "http://someweb.com/coolpage/and/not/cool/page"
	short := sh.Shorten(full)
	res := sh.Resolve(short)
	if res != full {
		t.Fatalf("got: %v, want: %v", res, full)
	}

	res = sh.Resolve("asdfwa")
	if res != "" {
		t.Fatalf("Expect an empty string, got '%v'", res)
	}
}
