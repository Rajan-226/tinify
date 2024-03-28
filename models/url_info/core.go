package url_info

type ICore interface {
}

type core struct {
	IRepo
}

func NewCore(repo IRepo) ICore {
	return &core{
		repo,
	}
}

func (c *core) NewEntity() {

}
