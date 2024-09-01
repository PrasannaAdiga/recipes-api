package main

import (
	"net/http"
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

func main() {
	router := gin.Default()
	router.POST("/receipe", NewReceipeHandler)
	router.Run()
}
