package endpoints

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/williamtran892/Shortener/service"
)

type GenerateUrlRequest struct {
	destinationURL string
}
type GenerateUrlResponse struct {
	srcURL string
}

type genUrlEndpoint struct {
	cfg string
	svc service.Linkage
}

func (e *genUrlEndpoint) DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// todo: read from body to populate destination URL
	return GenerateUrlRequest{
		destinationURL: "google.com",
	}, nil
}

func (e *genUrlEndpoint) ProcessRequest(ctx context.Context, r interface{}) (interface{}, error) {
	params := r.(GenerateUrlRequest)
	srcURL, err := e.svc.GenerateShortenedUrl(ctx, params)
	return GenerateUrlResponse{
		srcURL: srcURL,
	}, err
}

func (e *genUrlEndpoint) EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func newGenUrlEndpoint(conf string, svc service.Linkage) *httptransport.Server {
	// TODO: include middle
	ep := &genUrlEndpoint{cfg: conf, svc: svc}

	options := []httptransport.ServerOption{
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		// httptransport.ServerAfter(setResponseHeaders(ep)...),
		// httptransport.ServerErrorEncoder(newErrorEncoder(ep)),
		// httptransport.ServerErrorHandler(newErrorHandler()),
	}

	return httptransport.NewServer(
		ep.ProcessRequest,
		ep.DecodeRequest,
		ep.EncodeResponse,
		options...)
}
