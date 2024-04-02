package tinify

import (
	"context"
	"github.com/myProjects/tinify/internal/pkg/constants"
	"github.com/myProjects/tinify/internal/pkg/zookeeper"
	"github.com/myProjects/tinify/models/domain_info"
	"github.com/myProjects/tinify/models/url_info"
	"sync"
)

var (
	c    ICore
	once sync.Once
)

type ICore interface {
	GetShortened(context context.Context, url string) (string, error)
	GetLongURL(context context.Context, url string) (string, error)
	Tinify(context context.Context, url string, shorten strategy) (string, error)
	CreateAnalytics(context context.Context, url string) error
}

type core struct {
	urlCore    url_info.ICore
	domainCore domain_info.ICore
}

func NewCore(urlCore url_info.ICore, domainCore domain_info.ICore) {
	once.Do(
		func() {
			c = &core{
				urlCore:    urlCore,
				domainCore: domainCore,
			}
		},
	)
}

func (c *core) GetShortened(ctx context.Context, url string) (string, error) {
	urlInfo, err := c.urlCore.Fetch(ctx, url)
	if err != nil {
		return "", err
	}

	return urlInfo.GetShortenedURL(), nil
}

func (c *core) GetLongURL(ctx context.Context, url string) (string, error) {
	urlInfo, err := c.urlCore.Fetch(ctx, url)
	if err != nil {
		return "", err
	}

	return urlInfo.GetLongURL(), nil
}

func (c *core) Tinify(ctx context.Context, url string, shorten strategy) (string, error) {
	shortenedURI := shorten(zookeeper.GetCounter())

	if err := c.urlCore.NewEntity(ctx, url, shortenedURI); err != nil {
		return "", err
	}

	return constants.TinifyPrefixURL + shortenedURI, nil
}

func (c *core) CreateAnalytics(ctx context.Context, url string) error {
	return nil
}

func GetCore() ICore {
	return c
}
