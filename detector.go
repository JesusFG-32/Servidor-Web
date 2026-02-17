package main

import (
	"strings"
)

func detectDetails(ua string) (os, browser, bgColor, textColor string) {
	uaLower := strings.ToLower(ua)

	// Default values
	os = "Unknown"
	browser = "Unknown"
	bgColor = "#FFFFFF" // White
	textColor = "#000000"

	// OS Detection & Color
	if strings.Contains(uaLower, "windows") {
		os = "Windows"
		bgColor = "#0078D7" // Windows Blue
		textColor = "#FFFFFF"
	} else if strings.Contains(uaLower, "mac os") || strings.Contains(uaLower, "macintosh") {
		os = "macOS"
		bgColor = "#F0F0F0" // Mac Light Gray
		textColor = "#000000"
	} else if strings.Contains(uaLower, "android") {
		os = "Android"
		bgColor = "#3DDC84" // Android Green
		textColor = "#000000"
	} else if strings.Contains(uaLower, "linux") {
		os = "Linux"
		bgColor = "#333333" // Dark Gray for Linux
		textColor = "#FFFFFF"
	} else if strings.Contains(uaLower, "iphone") || strings.Contains(uaLower, "ipad") {
		os = "iOS"
		bgColor = "#000000" // Black for iOS
		textColor = "#FFFFFF"
	}

	// Browser Detection (Simple heuristic)
	if strings.Contains(uaLower, "firefox") {
		browser = "Mozilla Firefox"
	} else if strings.Contains(uaLower, "chrome") {
		browser = "Google Chrome"
	} else if strings.Contains(uaLower, "safari") && !strings.Contains(uaLower, "chrome") {
		browser = "Safari"
	} else if strings.Contains(uaLower, "edge") {
		browser = "Microsoft Edge"
	} else if strings.Contains(uaLower, "opera") || strings.Contains(uaLower, "opr") {
		browser = "Opera"
	} else {
		browser = "Unknown Browser"
	}

	return
}
