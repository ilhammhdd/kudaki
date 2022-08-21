package rest

import (
	"net/http"
)

type HandlerMiddleware struct {
	Handler http.Handler
}

type MethodRouting struct {
	PostHandler   http.Handler
	PutHandler    http.Handler
	DeleteHandler http.Handler
	GetHandler    http.Handler
	PatchHandler  http.Handler
}

func (mr MethodRouting) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.Method == http.MethodPost && mr.PostHandler != nil:
		mr.PostHandler.ServeHTTP(w, r)
	case r.Method == http.MethodPut && mr.PutHandler != nil:
		mr.PutHandler.ServeHTTP(w, r)
	case r.Method == http.MethodDelete && mr.DeleteHandler != nil:
		mr.DeleteHandler.ServeHTTP(w, r)
	case r.Method == http.MethodGet && mr.GetHandler != nil:
		mr.GetHandler.ServeHTTP(w, r)
	case r.Method == http.MethodPatch && mr.PatchHandler != nil:
		mr.PatchHandler.ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
