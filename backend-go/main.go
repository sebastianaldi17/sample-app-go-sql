package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/handler"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/repo"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/usecase"
)

var (
	jwtTokenAuth *jwtauth.JWTAuth
)

func main() {
	_, dbEnvPresent := os.LookupEnv("DB_STR")
	if !dbEnvPresent {
		log.Fatalln("DB_STR is empty, please check")
	}

	connectionString := os.Getenv("DB_STR")

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalln("JWT_SECRET is not set or empty, please check")
	}

	jwtTokenAuth = jwtauth.New("HS256", []byte("my-secret"), nil)

	repo := repo.New(db)
	usecase := usecase.New(*repo, jwtTokenAuth)
	handler := handler.New(*usecase)

	r := chi.NewRouter()

	// "public" routes
	r.Group(func(r chi.Router) {
		r.Get("/", handler.Hello)
		r.Post("/user", handler.CreateAccount)
		r.Post("/user/login", handler.LoginUser)
		r.Get("/todo", handler.GetTodos)
		r.Get("/todo/{todoID}", handler.GetTodoByID)
	})

	// protected routes
	r.Group(func(r chi.Router) {
		// Authentication middleware
		r.Use(jwtauth.Verifier(jwtTokenAuth))
		r.Use(jwtauth.Authenticator(jwtTokenAuth))

		r.Post("/todo", handler.InsertTodo)
		r.Put("/todo/{todoID}", handler.UpdateTodo)
		r.Delete("/todo/{todoID}", handler.DeleteTodo)
		r.Get("/user/todo", handler.GetTodosByUser)
	})

	log.Println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}
