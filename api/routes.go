package api

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

//go:embed static
var embeddedFS embed.FS

func Routes(cfg *Config) http.Handler {
	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Heartbeat("/health"))

	router.Get("/time", cfg.getTime)

	serverRoot, err := fs.Sub(embeddedFS, "static")
	if err != nil {
		log.Fatal(err)
	}

	router.Handle("/new/*", http.StripPrefix("/new", http.FileServer(http.FS(serverRoot))))

	router.Route("/{shortUrl}", func(r chi.Router) {
		r.Get("/", cfg.redirect)
	})

	router.Route("/shorten", func(r chi.Router) {
		r.Post("/", cfg.shorten)
	})

	router.Route("/", func(r chi.Router) {
		r.Handle("/", http.RedirectHandler("/new/", http.StatusPermanentRedirect))
	})

	return router
}
