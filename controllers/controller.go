package controllers

import (
	"log"
	"net/http"
	"task-organizer/models"
	"task-organizer/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service services.Service
}

func New() Controller {
	t := Controller{}
	t.Service = services.New()
	return t
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

func (t Controller) Signup(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	err := t.Service.CreateUser(user)
	if err != nil {
		log.Print(err)
	}
	c.Request.Method = http.MethodGet
	c.Redirect(http.StatusSeeOther, "/login")
}
