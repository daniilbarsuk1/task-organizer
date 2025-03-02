package main

import (
	"fmt"
	"net/http"
	"task-organizer/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (t Controller) Register(r *gin.Engine) {
	r.LoadHTMLGlob("views/**")
	r.Static("/public", "public")

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.POST("/signup", t.Signup)
}

func (r Controller) Signup(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	fmt.Print("data: " + user.Username)
}
