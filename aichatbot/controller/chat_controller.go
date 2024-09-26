package controller

import (
	"log"
	"net/http"
	"slack-chatbot/models"
	"slack-chatbot/requsts"

	"github.com/gin-gonic/gin"
)

var geminiAPIKey = "AIzaSyBUiRYimT0GT6ZndhR4_2CfD4UoYyzeVpo"

func GenerateHandler(c *gin.Context) {
	var prompt models.Prompt
	if err := c.BindJSON(&prompt); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	responseText, err := requsts.GetGeminiResponse(geminiAPIKey, prompt.Prompt)
	if err != nil {
		log.Printf("Error getting response from Gemini: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't process your request"})
		return
	}

	// Return the response back to the client
	c.JSON(http.StatusOK, gin.H{"response": responseText})
}
