package utils

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/viniciusbls9/go-movie/pkg/usecases"
)

func CreateRouters() error {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", usecases.GetMovies)
	v1Router.Get("/movies", usecases.GetMovies)
	v1Router.Post("/movies", usecases.CreateMovie)

	router.Mount("/v1", v1Router)

	fmt.Printf("Starting server at port 8000")
	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
