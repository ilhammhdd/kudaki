package rest

import (
	"net/http"
)

type EventDrivenHandler interface {
	http.Handler
	validate(r *http.Request) (errs *[]string, ok bool)
}
