package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type experience struct {
	ID               string `json:"id"`
	Company          string `json:"company"`
	Title            string `json:"title`
	DateFrom         string `json:"datefrom"`
	DateTo           string `json: "dateto"`
	Responsibilities string `json:"responsibilities"`
}

var allExperience = []experience{
	{
		ID:               "1",
		Company:          "HSBC",
		Title:            "Graduate Cloud Engineer",
		DateFrom:         "July 2024",
		DateTo:           "Present",
		Responsibilities: "XYZ",
	},
}

func main() {
	router := gin.Default()
	router.GET("/experience", getJobExperience)

	router.Run("localhost:8080")
}

// Get all experience
func getJobExperience(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allExperience)
}
