package model

import (
	"errors"
	"net/url"
)

// Page model
type Page struct {
	Domain string
	Path   string
}

// TODO: implement
func (p *Page) Validate() error {
	return nil
}

var (
	ErrFailedToParseURL = errors.New("failed to parse the url")
)

func NewPageFromURL(s string) (Page, error) {
	u, err := url.Parse(s)
	if err != nil {
		return Page{}, err
	}
	if len(u.Host) == 0 {
		u, err = url.Parse("http://" + s)
		if err != nil {
			return Page{}, err
		}
	}
	if len(u.Host) == 0 {
		return Page{}, ErrFailedToParseURL
	}
	return Page{
		Domain: u.Hostname(),
		Path:   u.Path,
	}, nil
}

// PageWithViews is a page with a total view count
type PageWithViews struct {
	Page
	Views int
}
