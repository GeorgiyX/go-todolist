package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-todolist/app/handlers"
	"go-todolist/app/repositories"
	"go-todolist/app/usecases"

	repoImpl "go-todolist/app/repositories/impl"
	ucImpl "go-todolist/app/usecases/impl"
)

type Repositories struct {
	Task repositories.ITaskRepository
}

type UseCases struct {
	Task usecases.ITaskUseCase
}

type Handlers struct {
	Task *handlers.TaskHandler
}

type Server struct {
	Settings     *Settings
	Repositories Repositories
	UseCases     UseCases
	Handlers     Handlers
}

func CreateServer() *Server {
	server := &Server{
		Settings:     LoadSettings(),
		Repositories: Repositories{},
		UseCases:     UseCases{},
		Handlers:     Handlers{},
	}

	return server
}

func (server *Server) Run() {
	/* DataBase */
	db, err := gorm.Open(postgres.Open(server.Settings.DSN), &gorm.Config{})
	if err != nil {
		fmt.Printf("Can't create DB connection: %v", err)
		return
	}

	/* Repositories & UseCases*/
	server.Repositories.Task = repoImpl.CreateTaskRepository(db)
	server.UseCases.Task = ucImpl.CreateTaskUseCase(server.Repositories.Task)

	/* Server */
	gin.SetMode(server.Settings.MODE)
	router := gin.New()
	if server.Settings.MODE == "debug" {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())
	apiGroup := router.Group(server.Settings.APIPrefix)

	/*Handlers*/
	server.Handlers.Task = handlers.CreateTaskHandler("/task", server.UseCases.Task, apiGroup)

	err = router.Run(server.Settings.APIAddr)
	if err != nil {
		fmt.Printf("Can't start server: %v\n", err)
		return
	}
}
