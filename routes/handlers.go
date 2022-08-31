package routes

import (
	"net/http"

	"timbul.io/service/backend"
)

type HandlerFunc func(*backend.Service, http.ResponseWriter, *http.Request)

func HandlerFuncWrapper(env *backend.Service, f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(env, w, r)
	}
}
