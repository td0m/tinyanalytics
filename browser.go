package model

// Browser enum
type Browser string

// Browsers
const (
	BrowserFirefox Browser = "firefox"
	BrowserChrome  Browser = "chrome"
	BrowserSafari  Browser = "safari"
	BrowserIE      Browser = "ie"
	BrowserUnknown Browser = "unknown"
)
