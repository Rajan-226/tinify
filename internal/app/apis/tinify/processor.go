package tinify

import (
	"context"
	"sync"
)

func Process(ctx context.Context, url string, core ICore) (string, error) {
	if present, shortURL := core.isAlreadyShortened(ctx, url); present {
		return shortURL, nil
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var shortURL string
	var err error

	go func() {
		shortURL, err = core.tinify(ctx, url, Base62Strategy)
	}()
	go func() {
		err = core.analytics(ctx, url)
	}()

	wg.Wait()

	if err != nil {
		return "", err
	}

	return shortURL, nil
}
