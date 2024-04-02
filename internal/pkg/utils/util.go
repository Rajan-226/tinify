package utils

import (
	urlParser "net/url"
)

func SplitURL(url string) (string, string, error) {
	parsedURL, err := urlParser.Parse(url)
	if err != nil {
		return "", "", err
	}

	domain := parsedURL.Hostname()
	uri := parsedURL.Path

	return domain, uri, nil
}

func IsValidURL(url string) bool {
	u, err := urlParser.Parse(url)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
