package shortener

type GenerateUrlRequest struct {
	destinationURL string
}
type GenerateUrlResponse struct {
	srcURL string
}

type ViewStatsRequest struct {
	SourceUrl string
}
type ViewStatsResponse struct {
	numOfVisited int
}
