package main

import (
	"finalproject/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.Static("/views", "./views")
	r.LoadHTMLGlob("views/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/tableqa", controllers.TableQA)
	r.POST("/summarize", controllers.Summarize)
	r.POST("/chatme", controllers.ChatMe)

	r.Run()
}
