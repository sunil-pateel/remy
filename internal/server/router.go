package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sunil-pateel/personal-website/cmd/templates"
)

func MakeRoutesHandler() *chi.Mux {
    r := chi.NewRouter()
    
    r.Use(middleware.Logger)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        var helloComponent templ.Component = templates.Hello("Sunil")  
        templ.Handler(helloComponent).ServeHTTP(w,r)
    })

    return r
}
