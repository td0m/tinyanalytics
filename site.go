package model

// Site model
type Site struct {
	Domain string `json:"domain,omitempty"`
	Owner  string `json:"owner,omitempty"`
}
