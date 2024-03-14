package main

import (
	"context"
	"fmt"
	"learngo/handlers"
	"learngo/settings"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	// these are constants because the db credentials should not change
	serverAddress string = ":8080"
	DB_HOST       string = "localhost"
	DB_PORT       int    = 5432
	DB_USERNAME   string = "root"
	DB_PASSWORD   string = "secret"
	DB_NAME       string = "learngo"
)

func main() {
	// make db connection
	dbSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME)
	dbClient, err := connectPostgres(dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Close()

	if dbClient == nil {
		log.Fatal("db connection failed")
	}

	// run api server
	runServer()
}

// connectPostgres makes db connection and return db client or error
func connectPostgres(source string) (*pgxpool.Pool, error) {
	var err error
	client, err := pgxpool.Connect(context.Background(), source)
	if err != nil {
		return nil, err
	}
	settings.DBClient = client

	if err = client.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Successfully connected with postgres db!")

	return client, nil
}

// runserver initates the go chi api server
func runServer() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	routes(router)

	log.Printf("server running on port %s", serverAddress)
	http.ListenAndServe(serverAddress, router)
}

// routes host all the application routes
func routes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to Go lang"))
	})
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("check"))
	})

	// group other routes with /api
	router.Route("/api", func(r chi.Router) {
		userRoutes(r)
		LoginRoutes(r)
	})
}

func userRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.UserCreateHandler)
		r.Get("/{id}", handlers.UserHandler)
		r.Get("/", handlers.UserListHandler)
	})
}

func LoginRoutes(r chi.Router) {
	r.Route("/login", func(r chi.Router) {
		r.Post("/", handlers.LoginHandler)
	})
}
