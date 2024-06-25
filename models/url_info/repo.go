package url_info

import (
	"context"
	"errors"

	"github.com/Rajan-226/tinify/internal/pkg/constants"
)

type IRepo interface {
	Create(ctx context.Context, key string, info *UrlInfo)
	Fetch(ctx context.Context, key string) (*UrlInfo, error)
	FetchAll(ctx context.Context) map[string]*UrlInfo
}

type repo struct {
	// we can add data source here, but as we need in memory, so will be directly using DB
	DB map[string]*UrlInfo
}

func NewRepo() IRepo {
	return &repo{
		DB: make(map[string]*UrlInfo), //ideally this should have been injected from upper level
	}
}

func (r *repo) Create(ctx context.Context, key string, info *UrlInfo) {
	r.DB[key] = info
}

func (r *repo) Fetch(ctx context.Context, key string) (*UrlInfo, error) {
	value, ok := r.DB[key]
	if !ok {
		return nil, errors.New(constants.NotFound)
	}

	if value == nil {
		return nil, errors.New("nil value")
	}

	return value, nil
}

func (r *repo) FetchAll(ctx context.Context) map[string]*UrlInfo {
	return r.DB
}
