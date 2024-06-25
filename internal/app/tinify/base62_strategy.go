package tinify

import (
	"github.com/Rajan-226/tinify/internal/pkg/constants"
)

type strategy func(counter int64) string

func Base62Strategy(counter int64) string {
	var result string

	if counter == 0 {
		return string(constants.Base62CharSet[0])
	}

	for counter > 0 {
		remainder := counter % 62
		result = string(constants.Base62CharSet[remainder]) + result
		counter /= 62
	}

	return result
}
