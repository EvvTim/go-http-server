package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}

	router := setupRouter()

	port := os.Getenv("PORT")
	if err := startServer(port, router); err != nil {
		log.Fatal(err)
	}
}

func loadConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	if os.Getenv("PORT") == "" {
		return &configError{"PORT is not found in the env"}
	}
	return nil
}

func setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*example.com", "http://*example.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)
	router.Mount("/v1", v1Router)

	return router
}

func startServer(port string, handler http.Handler) error {
	addr := ":" + port
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Printf("Server starting on port %v", addr)
	return srv.ListenAndServe()
}

type configError struct {
	msg string
}

func (e *configError) Error() string {
	return e.msg
}
