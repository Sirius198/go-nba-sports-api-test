package team

import (
	"encoding/json"
	"log"

	"github.com/Sirius198/go-nba-sports-api-test/consumer"
	"github.com/Sirius198/go-nba-sports-api-test/database"
)

func FetchAllTeams() {

	retry := 3
	db := database.DB

	for retry > 0 {

		resp, err := consumer.Get("/nba/scores/json/teams")
		if err != nil {
			retry--
			continue
		}

		var teams []Team
		if err := json.Unmarshal(resp, &teams); err != nil {
			retry--
			continue
		}

		var tmp Team
		for _, team := range teams {
			tmp.TeamID = team.TeamID
			tx := db.First(&tmp)
			if tx.RowsAffected == 0 {
				db.Create(&team)
			} else {
				db.Save(&team)
			}
		}

		log.Println("Successfully updated Teams to database")
		break
	}
}
