package service

import (
	"context"

	shortener "github.com/williamtran892/Shortener"
)

// Contains all services functions/interfaces responsible for keeping
// information about visited link

type ViewerStats interface {
	GetLinkageStats(ctx context.Context, params shortener.ViewStatsRequest) (string, error)
	// ResetLinkageStats(ctx context.Context, source string) error
}

type viewerStats struct{}

func NewViewerStatsService() *viewerStats {
	return &viewerStats{}
}

func (v *viewerStats) GenerateShortenedUrl(ctx context.Context, params shortener.ViewStatsRequest) (shortener.ViewStatsResponse, error) {
	return shortener.ViewStatsResponse{
		numOfVisited: 1,
	}, nil
}
