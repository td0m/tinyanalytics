package model

import (
	"reflect"
	"testing"
)

func TestNewPageFromURL(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Page
		wantErr bool
	}{
		{"parses basic page with protocol", args{"https://tdom.dev/hello"}, Page{Domain: "tdom.dev", Path: "/hello"}, false},
		{"parses basic page with port", args{"http://tdom.dev:69/hello"}, Page{Domain: "tdom.dev", Path: "/hello"}, false},
		{"parses basic page without protocol", args{"tdom.dev/hello"}, Page{Domain: "tdom.dev", Path: "/hello"}, false},
		{"parses subdomain", args{"dom.tdom.dev/hello"}, Page{Domain: "dom.tdom.dev", Path: "/hello"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPageFromURL(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPageFromURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPageFromURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
