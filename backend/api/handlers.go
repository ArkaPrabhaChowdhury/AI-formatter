package api

import (
	"gemini-backend/db"
	"gemini-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveData(c *gin.Context) {
	var input struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	_, err := db.DB.Exec("INSERT INTO data (content) VALUES ($1)", input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}

func FetchData(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, content FROM data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer rows.Close()

	var entries []db.DataEntry
	for rows.Next() {
		var entry db.DataEntry
		if err := rows.Scan(&entry.ID, &entry.Content); err != nil {
			continue
		}
		entries = append(entries, entry)
	}

	c.JSON(http.StatusOK, entries)
}

func FormatData(c *gin.Context) {
	var requestBody struct {
		Text       string `json:"text"`
		FormatType string `json:"formatType"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	formattedText, err := services.FormatDataWithGemini(requestBody.Text, requestBody.FormatType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"formattedText": formattedText})
}
