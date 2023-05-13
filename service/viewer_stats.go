package service

import (
	"context"

	"github.com/williamtran892/Shortener/endpoints"
)

// Contains all services functions/interfaces responsible for keeping
// information about visited link

type ViewerStats interface {
	GetLinkageStats(ctx context.Context, params endpoints.ViewStatsRequest) (string, error)
	// ResetLinkageStats(ctx context.Context, source string) error
}

type viewerStats struct{}

func NewViewerStatsService() *viewerStats {
	return &viewerStats{}
}

func (v *viewerStats) GenerateShortenedUrl(ctx context.Context, params endpoints.ViewStatsRequest) (endpoints.ViewStatsResponse, error) {
	return endpoints.ViewStatsResponse{
		numOfVisited: 1,
	}, nil
}
