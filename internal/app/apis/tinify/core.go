package tinify

import (
	"github.com/myProjects/tinify/models/domain_info"
	"github.com/myProjects/tinify/models/url_info"
)

type ICore interface {
}

type core struct {
	urlCore    url_info.ICore
	domainCore domain_info.ICore
}

func NewCore() ICore {
	return &core{}
}
