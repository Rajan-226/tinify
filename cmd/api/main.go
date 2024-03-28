package main

import (
	"context"
	"github.com/myProjects/tinify/internal/pkg/boot"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	boot.Init(ctx)
}

/*

url -> shortUrlMapping

shortURLGenerating
shortURL Redirection


*/
