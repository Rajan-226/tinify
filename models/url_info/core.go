package url_info

import (
	"context"
	"time"
)

type ICore interface {
	NewEntity(ctx context.Context, url string, shortURI string) error
	Fetch(ctx context.Context, longURL string) (*UrlInfo, error)
}

type core struct {
	repo IRepo
}

func NewCore(repo IRepo) ICore {
	return &core{
		repo: repo,
	}
}

func (c *core) NewEntity(ctx context.Context, longURL string, shortURI string) error {
	createdAt := time.Now().UnixMilli()
	entity := &UrlInfo{
		ShortURI:   shortURI,
		LongURL:    longURL,
		VisitCount: 0,
		Expiry:     createdAt + int64(5*24*time.Hour),
		CreatedAt:  createdAt,
	}

	c.repo.Create(ctx, longURL, entity)

	return nil
}

func (c *core) Fetch(ctx context.Context, longURL string) (*UrlInfo, error) {
	entity, err := c.repo.Fetch(ctx, longURL)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
