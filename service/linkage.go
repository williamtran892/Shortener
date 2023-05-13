package service

// Contains all services functions/interfaces responsible for
// managing the mapping source to destination URLs
// (hence the name linkage)

import (
	"context"
)

type Linkage interface {
	GenerateShortenedUrl(ctx context.Context, destination string) (string, error)
	// GenerateShortenedUrlWithSpecificSrc(ctx context.Context, source, destination string) error
	// UpdateDestinationURL(ctx context.Context, source, destination string) error
	// DisableLinkage(ctx context.Context, source string) error
}

type linkage struct{}

func NewService() *linkage {
	return &linkage{}
}

func (l *linkage) GenerateShortenedUrl(ctx context.Context, destination string) (string, error) {
	return "newShortenedPath", nil
}
