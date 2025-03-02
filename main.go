package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	c := Controller{}
	c.Register(router)
	router.Run(":8080")
}
