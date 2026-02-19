package main

import (
	"strings"
)

func detectDetails(ua string) (os, browser, bgColor, textColor string) {
	uaLower := strings.ToLower(ua)
	os = "Unknown"
	browser = "Unknown"
	bgColor = "#FFFFFF"
	textColor = "#000000"

	if strings.Contains(uaLower, "windows") {
		os = "Windows"
		bgColor = "#0078D7"
		textColor = "#FFFFFF"
	} else if strings.Contains(uaLower, "mac os") || strings.Contains(uaLower, "macintosh") {
		os = "macOS"
		bgColor = "#F0F0F0"
		textColor = "#000000"
	} else if strings.Contains(uaLower, "android") {
		os = "Android"
		bgColor = "#3DDC84"
		textColor = "#000000"
	} else if strings.Contains(uaLower, "iphone") || strings.Contains(uaLower, "ipad") {
		os = "iOS"
		bgColor = "#000000"
		textColor = "#FFFFFF"
	} else if strings.Contains(uaLower, "linux") {
		if strings.Contains(uaLower, "ubuntu") {
			os = "Ubuntu"
			bgColor = "#E95420"
			textColor = "#FFFFFF"
		} else if strings.Contains(uaLower, "debian") {
			os = "Debian"
			bgColor = "#A80030"
			textColor = "#FFFFFF"
		} else if strings.Contains(uaLower, "fedora") {
			os = "Fedora"
			bgColor = "#294172"
			textColor = "#FFFFFF"
		} else if strings.Contains(uaLower, "arch") {
			os = "Arch Linux"
			bgColor = "#1793D1"
			textColor = "#FFFFFF"
		} else if strings.Contains(uaLower, "mint") {
			os = "Linux Mint"
			bgColor = "#87CF3E"
			textColor = "#000000"
		} else if strings.Contains(uaLower, "gentoo") {
			os = "Gentoo"
			bgColor = "#54487A"
			textColor = "#FFFFFF"
		} else {
			os = "Linux"
			bgColor = "#333333"
			textColor = "#FFFFFF"
		}
	}

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
