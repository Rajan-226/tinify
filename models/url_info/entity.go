package url_info

import "github.com/Rajan-226/tinify/internal/pkg/constants"

type UrlInfo struct {
	LongURL    string
	ShortURI   string
	VisitCount int64
	Expiry     int64
	CreatedAt  int64
}

func (u *UrlInfo) GetShortenedURL() string {
	return constants.TinifyPrefixURL + u.ShortURI
}

func (u *UrlInfo) GetLongURL() string {
	return u.LongURL
}
