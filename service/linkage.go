package service

// Contains all services functions/interfaces responsible for
// managing the mapping source to destination URLs
// (hence the name linkage)

import (
	"context"

	"github.com/williamtran892/Shortener/endpoints"
)

type Linkage interface {
	GenerateShortenedUrl(ctx context.Context, params endpoints.GenerateUrlRequest) (string, error)
	// GenerateShortenedUrlWithSpecificSrc(ctx context.Context, source, destination string) error
	// UpdateDestinationURL(ctx context.Context, source, destination string) error
	// DisableLinkage(ctx context.Context, source string) error
}

type linkage struct{}

func NewLinkageService() *linkage {
	return &linkage{}
}

func (l *linkage) GenerateShortenedUrl(ctx context.Context, destination string) (string, error) {
	return "newShortenedPath", nil
}
