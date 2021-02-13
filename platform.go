package model

// Platform enum
type Platform string

// Platforms
const (
	PlatformWindows  Platform = "windows"
	PlatformMac      Platform = "mac"
	PlatformLinux    Platform = "linux"
	PlatformIPad     Platform = "ipad"
	PlatformIPhone   Platform = "iphone"
	PlatformAndroid  Platform = "android"
	PlatformChromeOS Platform = "chromeos"
	PlatformUnknown  Platform = "unknown"
)
