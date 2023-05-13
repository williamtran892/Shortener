package endpoints

import (
	"context"
	"net/http"
)

type GenerateUrlRequest string

type genUrlEndpoint struct {
	cfg string
	svc service.Linkage
}

func (e *genUrlEndpoint) DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (e *genUrlEndpoint) ProcessRequest(ctx context.Context, r interface{}) (interface{}, error) {
	params := r.(GenerateUrlRequest)
	err := e.service.DeleteGroup(ctx, params)
	return nil, err
}

func (e *genUrlEndpoint) EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
	return nil
}
