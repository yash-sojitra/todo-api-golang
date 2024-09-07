package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yash-sojitra/todo/api/handlers"
	"github.com/yash-sojitra/todo/api/middleware"
	"github.com/yash-sojitra/todo/repository"
	"github.com/yash-sojitra/todo/services"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.Use(middleware.PathLogger)
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	taskRepo := repository.NewTaskRepository(s.db)
	taskService := services.NewTaskService(*taskRepo)
	taskHandler := handlers.NewTaskHandler(*taskService)
	taskHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}