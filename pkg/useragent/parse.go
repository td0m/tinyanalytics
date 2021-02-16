package useragent

import (
	"github.com/avct/uasurfer"
	model "github.com/td0m/tinyanalytics"
)

// Parser represents the user_agent parser
type Parser struct {
}

// NewParser creates a new user_agent parser
func NewParser() *Parser { return &Parser{} }

// Parse parses a user agent string
func (p *Parser) Parse(s string) (model.Platform, model.Browser) {
	userAgent := uasurfer.Parse(s)
	platform, browser := model.PlatformUnknown, model.BrowserUnknown
	switch userAgent.Browser.Name {
	case uasurfer.BrowserFirefox:
		browser = model.BrowserFirefox
	case uasurfer.BrowserChrome:
		browser = model.BrowserChrome
	case uasurfer.BrowserSafari:
		browser = model.BrowserSafari
	case uasurfer.BrowserIE:
		browser = model.BrowserIE
	default:
	}
	switch userAgent.OS.Platform {
	case uasurfer.PlatformWindows:
		platform = model.PlatformWindows
	case uasurfer.PlatformMac:
		platform = model.PlatformMac
	case uasurfer.PlatformLinux:
		platform = model.PlatformLinux
	case uasurfer.PlatformiPad:
		platform = model.PlatformIPad
	case uasurfer.PlatformiPhone:
		platform = model.PlatformIPhone
	}
	if userAgent.OS.Name == uasurfer.OSChromeOS {
		platform = model.PlatformChromeOS
	} else if userAgent.OS.Name == uasurfer.OSAndroid {
		platform = model.PlatformAndroid
	}

	return platform, browser
}
