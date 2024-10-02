package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"

	"github.com/sebastianaldi17/sample-app-go-sql/internal/handler"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/pkg/logger"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/repo"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/usecase"
)

var (
	jwtTokenAuth *jwtauth.JWTAuth
)

func main() {
	_, dbEnvPresent := os.LookupEnv("DB_STR")
	if !dbEnvPresent {
		panic("DB_STR is empty, please check")
	}

	connectionString := os.Getenv("DB_STR")

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	var app *newrelic.Application
	nrLicense := os.Getenv("NR_LICENSE")
	if len(nrLicense) == 40 { // NewRelic license length is 40
		app, err = newrelic.NewApplication(
			newrelic.ConfigAppName("sample-todo-app"),
			newrelic.ConfigLicense(os.Getenv("NR_LICENSE")),
			newrelic.ConfigAppLogForwardingEnabled(true),
		)
		if err != nil {
			panic(err)
		}
	}
	logger.InitLogging(app)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET is not set or empty, please check")
	}

	jwtTokenAuth = jwtauth.New("HS256", []byte("my-secret"), nil)

	repo := repo.New(db, redisClient)
	usecase := usecase.New(*repo, jwtTokenAuth)
	handler := handler.New(*usecase)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// "public" routes
	r.Group(func(r chi.Router) {
		if app != nil {
			r.Get(newrelic.WrapHandleFunc(app, "/", handler.Hello))
			r.Post(newrelic.WrapHandleFunc(app, "/user", handler.CreateAccount))
			r.Post(newrelic.WrapHandleFunc(app, "/user/login", handler.LoginUser))
		} else {
			r.Get("/", handler.Hello)
			r.Post("/user", handler.CreateAccount)
			r.Post("/user/login", handler.LoginUser)
		}
	})

	// protected routes
	r.Group(func(r chi.Router) {
		// Authentication middleware
		r.Use(jwtauth.Verifier(jwtTokenAuth))
		r.Use(jwtauth.Authenticator(jwtTokenAuth))

		if app != nil {
			r.Get(newrelic.WrapHandleFunc(app, "/user", handler.ValidateJWT))
			r.Post(newrelic.WrapHandleFunc(app, "/todo", handler.InsertTodo))
			r.Put(newrelic.WrapHandleFunc(app, "/todo/{todoID}", handler.UpdateTodo))
			r.Get(newrelic.WrapHandleFunc(app, "/todo/{todoID}", handler.GetTodoByID))
			r.Delete(newrelic.WrapHandleFunc(app, "/todo/{todoID}", handler.DeleteTodo))
			r.Get(newrelic.WrapHandleFunc(app, "/user/todo", handler.GetTodosByUser))

		} else {
			r.Get("/user", handler.ValidateJWT)
			r.Post("/todo", handler.InsertTodo)
			r.Put("/todo/{todoID}", handler.UpdateTodo)
			r.Get("/todo/{todoID}", handler.GetTodoByID)
			r.Delete("/todo/{todoID}", handler.DeleteTodo)
			r.Get("/user/todo", handler.GetTodosByUser)
		}
	})

	logger.Info("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}
