package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"timbul.io/controller"
	"timbul.io/lib/mongodb"
	"timbul.io/service/backend"
)

func CreateRoute() *chi.Mux {
	r := chi.NewRouter()

	db := mongodb.New()
	env := backend.New(&db)

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/movies", func(r chi.Router) {
		r.Get("/{movieId}", HandlerFuncWrapper(&env, controller.GetMovie))
	})

	return r
}
