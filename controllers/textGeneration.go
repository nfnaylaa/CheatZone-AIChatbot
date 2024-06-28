package controllers

import (
	"finalproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatMe(c *gin.Context) {
	message := c.PostForm("message")
	res, err := models.TextGeneration(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": *res})
}
