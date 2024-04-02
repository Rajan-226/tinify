package tinify

import (
	"context"
	"errors"
	"github.com/myProjects/tinify/internal/pkg/constants"
	"sync"
)

func Create(ctx context.Context, url string, core ICore) (string, error) {
	if shortURL, err := core.GetShortened(ctx, url); shortURL != "" {
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

	analyticsError := core.CreateAnalytics(ctx, url)

	wg.Wait()

	if analyticsError != nil {
		return "", analyticsError
	} else if err != nil {
		return "", err
	}

	return shortURL, nil
}

func Redirect(ctx context.Context, url string, core ICore) (string, error) {
	if longURL, err := core.GetLongURL(ctx, url); longURL != "" {
		return longURL, nil
	} else if err != nil && err.Error() != constants.NotFound {
		return "", err
	}

	return "", errors.New(constants.NotFound)
}
