package app

import (
	"log"
	"net/http"

	v1 "github.com/airats/bug_and_fix/internal/api/http/v1"
	api_v1 "github.com/airats/bug_and_fix/pkg/api/rest"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	serverImpl := v1.NewBugAndFixV1()

	strictHandler := api_v1.NewStrictHandler(serverImpl, nil)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	api_v1.HandlerFromMux(strictHandler, router)

	httpServer := &http.Server{
		Addr:    ":3000", // TODO: read from config
		Handler: router,
	}

	// TODO: gracful shutdown
	log.Printf("Starting server on %s", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
