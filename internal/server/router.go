package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sunil-pateel/personal-website/internal/server/routes"
	"github.com/sunil-pateel/personal-website/web/templates"
)

func MakeRoutesHandler() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

    fs := http.FileServer(http.Dir("./web"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	r.Mount("/search", routes.MakeSearchRouter())
	r.Get("/", IndexHandler)

	return r
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(templates.Index()).ServeHTTP(w, r)
}
