package main

import (
	"task-organizer/controllers"

	"github.com/gin-gonic/gin"
)

type App struct {
	controller controllers.Controller
}

func (a App) Start() {
	var router = gin.Default()
	a.controller = controllers.New()
	a.controller.Register(router)
	router.Run(":8080")
}

func main() {
	App{}.Start()
}
