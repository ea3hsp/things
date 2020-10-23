package http

import (
	"context"
	"net/http"

	"github.com/ea3hsp/test/api"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer ...
func NewHTTPServer(ctx context.Context, endpoints api.Endpoints) http.Handler {
	// goriall router instance
	r := mux.NewRouter()
	// common middle ware
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/things").Handler(httptransport.NewServer(
		endpoints.CreateThingEP,
		api.CreateThingsRequest,
		api.CreateThingsResponse,
	))
	r.Methods("GET").Path("/things/{key}").Handler(httptransport.NewServer(
		endpoints.ReadThingEP,
		api.ReadThingRequest,
		api.ReadThingResponse,
	))
	r.Methods("DELETE").Path("/things").Handler(httptransport.NewServer(
		endpoints.DeleteThingEP,
		nil,
		nil,
	))
	r.Methods("PUT").Path("/things").Handler(httptransport.NewServer(
		endpoints.UpdateThingEP,
		api.UpdateThingRequest,
		api.UpdateThingResponse,
	))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
