package url_info

var DB map[string]UrlInfo

type IRepo interface {
}

type repo struct {
	// we can add data source here, but as we need in memory, so will be directly using DB
}

func NewRepo() IRepo {
	return &repo{}
}
