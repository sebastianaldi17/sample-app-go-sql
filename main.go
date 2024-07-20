package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/entity"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/handler"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/repo"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/usecase"
	"gopkg.in/yaml.v3"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	configBytes, err := os.ReadFile(fmt.Sprintf("%s/files/%s.config.yaml", path, env))
	if err != nil {
		log.Fatalln(err)
	}

	var yamlConfig entity.Config
	err = yaml.Unmarshal(configBytes, &yamlConfig)
	if err != nil {
		log.Fatalln(err)
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		yamlConfig.Database.User,
		yamlConfig.Database.Password,
		yamlConfig.Database.Host,
		yamlConfig.Database.DatabaseName)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	repo := repo.New(db)
	usecase := usecase.New(*repo)
	handler := handler.New(*usecase)

	r := chi.NewRouter()

	r.Get("/todo", handler.GetTodos)
	r.Post("/todo", handler.InsertTodo)
	r.Get("/todo/{todoID}", handler.GetTodoByID)
	r.Delete("/todo/{todoID}", handler.DeleteTodo)
	r.Put("/todo/{todoID}", handler.UpdateTodo)
	r.Get("/", handler.Hello)

	log.Println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}
