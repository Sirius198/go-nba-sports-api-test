package game

import (
	"encoding/json"

	"github.com/Sirius198/go-nba-sports-api-test/consumer"
	"github.com/gin-gonic/gin"
)

func GetGamesByDate(c *gin.Context) {
	date := c.Param("date")

	// Process date to parse any format for date

	resp, err := consumer.Get("/nba/scores/json/GamesByDate/" + date)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	var games []Game
	if err := json.Unmarshal(resp, &games); err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	c.JSON(200, games)
}
