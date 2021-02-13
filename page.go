package model

// Page model
type Page struct {
	Domain string
	Path   string
}

// PageWithViews is a page with a total view count
type PageWithViews struct {
	Page
	Views int
}
