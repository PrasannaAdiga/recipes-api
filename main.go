package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid" //Global unique ID generator plugin
)

type Receipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var receipes []Receipe

func NewReceipeHandler(c *gin.Context) {
	var receipe Receipe
	if err := c.ShouldBindJSON(&receipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	receipe.ID = xid.New().String()
	receipe.PublishedAt = time.Now()
	receipes = append(receipes, receipe)
	c.JSON(http.StatusOK, receipe)
}

func ListReceipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, receipes)
}

func init() {
	file, _ := os.ReadFile("receipes.json")
	json.Unmarshal([]byte(file), &receipes)
}

func main() {
	router := gin.Default()
	router.POST("/receipes", NewReceipeHandler)
	router.GET("/receipes", ListReceipesHandler)
	router.Run()
}
