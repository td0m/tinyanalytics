package site

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const path = "tinyanalytics.txt"

// errors
var (
	ErrKeyFileNotFound = fmt.Errorf("Key file '%s' not found", path)
)

// FileKeyCrawler represents the site key crawler
type FileKeyCrawler struct {
}

// NewFileKeyCrawler creates a new site key crawler
func NewFileKeyCrawler() *FileKeyCrawler { return &FileKeyCrawler{} }

func (d *FileKeyCrawler) GetKey(domain string) (key string, err error) {
	resp, err := http.Get("http://" + domain + "/" + path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	s := strings.TrimSpace(string(bytes))
	if len(s) == 0 {
		return "", ErrKeyFileNotFound
	}
	return s, nil
}
