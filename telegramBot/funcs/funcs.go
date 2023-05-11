package funcs

import (
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func IsImageURL(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}
	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "image/") {
		return true
	}
	ext := strings.ToLower(filepath.Ext(url))
	if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp" {
		return true
	}
	return false
}

func IsURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	u, err := url.Parse(input)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
