package controllers

import (
	"finalproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Summarize(c *gin.Context) {
	input := c.PostForm("input")
	res, err := models.Summary(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": *res})
}
