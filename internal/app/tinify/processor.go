package tinify

import (
	"context"
	"errors"
	"sync"

	"github.com/Rajan-226/tinify/internal/pkg/constants"
)

func Create(ctx context.Context, url string, core ICore) (string, error) {
	if shortURL, err := core.GetShortenedURLIfExists(ctx, url); shortURL != "" {
		return shortURL, nil
	} else if err != nil && err.Error() != constants.NotFound {
		return "", err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	var shortURL string
	var err error

	go func() {
		defer wg.Done()
		shortURL, err = core.Tinify(ctx, url, Base62Strategy)
	}()

	analyticsError := core.Analytics(ctx, url)

	wg.Wait()

	if analyticsError != nil {
		return "", analyticsError
	} else if err != nil {
		return "", err
	}

	return shortURL, nil
}

func GetAllUrls(ctx context.Context, core ICore) (map[string]string, error) {
	return core.GetAllURLs(ctx)
}

func Redirect(ctx context.Context, url string, core ICore) (string, error) {
	if longURL, err := core.GetLongURL(ctx, url); longURL != "" {
		return longURL, nil
	} else if err != nil && err.Error() != constants.NotFound {
		return "", err
	}

	return "", errors.New(constants.NotFound)
}

func Metrics(ctx context.Context, core ICore) (map[string]int, error) {
	return core.GetTopShortenedDomains(ctx, 3)
}
