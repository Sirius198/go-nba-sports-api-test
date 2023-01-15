package player

import (
	"strconv"

	"github.com/Sirius198/go-nba-sports-api-test/database"
	"github.com/gin-gonic/gin"
)

func GetPlayerById(c *gin.Context) {
	playerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	_ = playerId
}

func GetPlayerDetailsByName(c *gin.Context) {
	firstName, lastName := c.Param("firstname"), c.Param("lastname")
	db := database.DB

	var p Player
	tx := db.Where(&Player{FirstName: firstName, LastName: lastName}).First(&p)
	if tx.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"status": false,
			"msg":    "Not found",
		})
	} else {
		c.JSON(200, gin.H{
			"status": true,
			"data":   p,
		})
	}
}

func GetPlayersByTeam(c *gin.Context) {
	team := c.Param("team")
	db := database.DB

	var players []Player
	db.Where(&Player{Team: team}).Find(&players)

	c.JSON(200, gin.H{
		"status": true,
		"data":   players,
	})
}
