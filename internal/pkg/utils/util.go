package utils

import "net/url"

func SplitURL(longURL string) (string, string, error) {
	parsedURL, err := url.Parse(longURL)
	if err != nil {
		return "", "", err
	}

	domain := parsedURL.Hostname()
	uri := parsedURL.Path

	return domain, uri, nil
}
