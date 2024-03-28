package url_info

type ICore interface {
}

type core struct{}

func NewCore() ICore {
	return &core{}
}

func (c *core) NewEntity() {

}
