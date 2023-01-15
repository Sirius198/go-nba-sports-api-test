package nba

// func GetGamesByDate(c *gin.Context) {
// 	var req GamesByDateRequest
// 	if err := c.BindUri(&req); err != nil {
// 		c.JSON(400, gin.H{"msg": err})
// 		return
// 	}

// 	resp, err := consumer.Get("/nba/scores/json/GamesByDate/" + req.Date)
// 	if err != nil {
// 		c.JSON(400, gin.H{"msg": err})
// 		return
// 	}

// 	var games []Game
// 	if err := json.Unmarshal(resp, &games); err != nil {
// 		c.JSON(500, gin.H{"msg": err})
// 		return
// 	}

// 	c.JSON(200, games)
// }

// func GetPlayersByTeam(c *gin.Context) {
// 	var req PlayersByTeamRequest
// 	if err := c.BindUri(&req); err != nil {
// 		c.JSON(400, gin.H{"msg": err})
// 		return
// 	}

// 	resp, err := consumer.Get("/nba/scores/json/Players/" + req.TeamName)
// 	if err != nil {
// 		c.JSON(400, gin.H{"msg": err})
// 		return
// 	}

// 	var players []Player
// 	if err := json.Unmarshal(resp, &players); err != nil {
// 		c.JSON(500, gin.H{"msg": err})
// 		return
// 	}

// 	c.JSON(200, players)
// }

// func GetPlayerDetailsByName(c *gin.Context) {
// 	var req PlayerDetailsByNameRequest
// 	if err := c.BindUri(&req); err != nil {
// 		c.JSON(400, gin.H{"msg": err})
// 		return
// 	}

// 	resp, err := consumer.Get("/nba/scores/json/Players")
// 	if err != nil {
// 		c.JSON(400, gin.H{"msg": err})
// 		return
// 	}

// 	var players []Player
// 	if err := json.Unmarshal(resp, &players); err != nil {
// 		c.JSON(500, gin.H{"msg": err})
// 		return
// 	}

// 	for _, player := range players {
// 		if player.FirstName == req.FirstName && player.LastName == req.LastName {
// 			c.JSON(200, gin.H{
// 				"status":   "success",
// 				"headshot": player.PhotoUrl,
// 			})
// 			return
// 		}
// 	}

// 	c.JSON(200, gin.H{
// 		"status":      "failed",
// 		"description": "Player not found",
// 	})
// }
