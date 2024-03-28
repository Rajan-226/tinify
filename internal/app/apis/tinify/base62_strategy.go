package tinify

type IStrategy func(longURL string) string

func Base62Strategy(longURL string) string {
	return ""
}
