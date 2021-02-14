package model

// Page model
type Page struct {
	Domain string
	Path   string
}

// TODO: implement
func (p *Page) Validate() error {
	return nil
}

// PageWithViews is a page with a total view count
type PageWithViews struct {
	Page
	Views int
}
