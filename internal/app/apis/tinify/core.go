package tinify

import (
	"github.com/myProjects/tinify/models/domain_info"
	"github.com/myProjects/tinify/models/url_info"
	"sync"
)

var (
	c    ICore
	once sync.Once
)

type ICore interface {
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

func GetCore() ICore {
	return c
}
