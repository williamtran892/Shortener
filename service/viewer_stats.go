package service

import "context"

// Contains all services functions/interfaces responsible for keeping
// information about visited link

type ViewerStats interface {
	GetLinkageStats(ctx context.Context, source string) (string, error)
	// ResetLinkageStats(ctx context.Context, source string) error
}

type viewerStats struct{}

func NewService() *viewerStats {
	return &viewerStats{}
}

func (v *viewerStats) GenerateShortenedUrl(ctx context.Context, destination string) (string, error) {
	return "nums_of_visited: 10", nil
}
