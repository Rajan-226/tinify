package tinify

import (
	"context"
	"github.com/myProjects/tinify/internal/pkg/utils"
	"github.com/redis/go-redis/v9"
	"sync"

	"github.com/myProjects/tinify/internal/pkg/constants"
	"github.com/myProjects/tinify/internal/pkg/zookeeper"
	"github.com/myProjects/tinify/models/url_info"
)

var (
	c    ICore
	once sync.Once
)

type ICore interface {
	GetShortened(context context.Context, url string) (string, error)
	GetLongURL(context context.Context, url string) (string, error)
	Tinify(context context.Context, url string, shorten strategy) (string, error)
	Analytics(context context.Context, url string) error
}

type core struct {
	urlCore url_info.ICore
	redis   redis.UniversalClient
}

func NewCore(urlCore url_info.ICore, client redis.UniversalClient) {
	once.Do(
		func() {
			c = &core{
				urlCore: urlCore,
				redis:   client,
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

func (c *core) Analytics(ctx context.Context, url string) error {
	domain, _, err := utils.SplitURL(url)
	if err != nil {
		return err
	}

	c.redis.Incr(ctx, domain)

	return nil
}

func GetCore() ICore {
	return c
}
