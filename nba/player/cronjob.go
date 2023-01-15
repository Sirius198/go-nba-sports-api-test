package player

import (
	"encoding/json"
	"log"

	"github.com/Sirius198/go-nba-sports-api-test/consumer"
	"github.com/Sirius198/go-nba-sports-api-test/database"
)

func FetchAllPlayers() {
	retry := 3
	db := database.DB

	for retry > 0 {

		resp, err := consumer.Get("/nba/scores/json/Players")
		if err != nil {
			retry--
			continue
		}

		var players []Player
		if err := json.Unmarshal(resp, &players); err != nil {
			retry--
			continue
		}

		var tmp Player
		for _, player := range players {
			tmp.PlayerID = player.PlayerID
			tx := db.First(&tmp)
			if tx.RowsAffected == 0 {
				db.Create(&player)
			} else {
				db.Save(&player)
			}
		}

		log.Println("Successfully updated Players to database")
		break
	}
}
