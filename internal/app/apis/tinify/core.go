package tinify

import (
	"context"
	"errors"
	"github.com/myProjects/tinify/models/domain_info"
	"github.com/myProjects/tinify/models/url_info"
	"sync"
)

var (
	c    ICore
	once sync.Once
)

type ICore interface {
	isAlreadyShortened(context context.Context, url string) (bool, string)
	tinify(context context.Context, url string, strategy IStrategy) (string, error)
	analytics(context context.Context, url string) error
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

func (c *core) isAlreadyShortened(ctx context.Context, url string) (bool, string) {
	return false, ""
}

func (c *core) tinify(ctx context.Context, url string, strategy IStrategy) (string, error) {

	return "", errors.New("tinify method not implemented")
}

func (c *core) analytics(ctx context.Context, url string) error {
	return nil
}

func GetCore() ICore {
	return c
}
