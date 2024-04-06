package url_info

import (
	"context"
	"errors"
	"github.com/myProjects/tinify/internal/pkg/utils"
	"time"
)

type ICore interface {
	NewEntity(ctx context.Context, url string, shortURI string) error
	Fetch(ctx context.Context, longURL string) (*UrlInfo, error)
	FetchAll(ctx context.Context) (map[string]string, error)
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
	c.repo.Create(ctx, shortURI, entity)

	return nil
}

func (c *core) FetchAll(ctx context.Context) (map[string]string, error) {
	entities := c.repo.FetchAll(ctx)

	longToShortURL := make(map[string]string)
	for key, value := range entities {
		if value == nil {
			return nil, errors.New("not a valid scenario")
		}

		if utils.IsValidURL(key) {
			longToShortURL[key] = value.ShortURI
		}
	}

	return longToShortURL, nil
}

func (c *core) Fetch(ctx context.Context, longURL string) (*UrlInfo, error) {
	entity, err := c.repo.Fetch(ctx, longURL)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
