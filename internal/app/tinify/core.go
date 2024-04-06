package tinify

import (
	"context"
	"fmt"
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
	GetShortenedURLIfExists(context context.Context, url string) (string, error)
	GetLongURL(context context.Context, url string) (string, error)
	GetAllURLs(context context.Context) (map[string]string, error)
	Tinify(context context.Context, url string, shorten strategy) (string, error)
	Analytics(context context.Context, url string) error
	GetTopShortenedDomains(ctx context.Context, count int64) (map[string]int, error)
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

func (c *core) GetShortenedURLIfExists(ctx context.Context, url string) (string, error) {
	urlInfo, err := c.urlCore.Fetch(ctx, url)
	if err != nil {
		return "", err
	}

	return urlInfo.GetShortenedURL(), nil
}

func (c *core) GetAllURLs(ctx context.Context) (map[string]string, error) {
	return c.urlCore.FetchAll(ctx)
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
	_, err = c.redis.ZIncrBy(ctx, "domains", 1, domain).Result()
	if err != nil {
		return fmt.Errorf("failed to increment domain count: %w", err)
	}

	return nil
}

func (c *core) GetTopShortenedDomains(ctx context.Context, count int64) (map[string]int, error) {
	domains, err := c.redis.ZRevRangeByScore(ctx, "domains", &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  count,
	}).Result()

	if err != nil {
		return nil, fmt.Errorf("error retrieving top domains: %w", err)
	}

	topDomains := make(map[string]int)
	for _, domain := range domains {
		cnt, err := c.redis.ZScore(ctx, "domains", domain).Result()
		if err != nil {
			return nil, fmt.Errorf("error retrieving count for domain %s: %w", domain, err)
		}
		topDomains[domain] = int(cnt)
	}

	return topDomains, nil
}

func GetCore() ICore {
	return c
}
