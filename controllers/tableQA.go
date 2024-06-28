package controllers

import (
	"io"
	"log"
	"net/http"

	"finalproject/models"

	"github.com/gin-gonic/gin"
)

func TableQA(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error getting form file:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	query := c.PostForm("query")
	log.Println("Query received:", query)

	f, err := file.Open()
	if err != nil {
		log.Println("Error opening file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer f.Close()

	fileContent, err := io.ReadAll(f)
	if err != nil {
		log.Println("Error reading file content:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	data := string(fileContent)
	log.Println("CSV data received:", data)

	table := csvToMap(data)
	log.Println("CSV data converted to map:", table)

	res, err := models.TableQA(query, table)
	if err != nil {
		log.Println("Error in TableQA model:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": res.Answer})
}
