package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"

	nbadatabase "github.com/Sirius198/go-nba-sports-api-test/database"
	nbagame "github.com/Sirius198/go-nba-sports-api-test/nba/game"
	nbaplayer "github.com/Sirius198/go-nba-sports-api-test/nba/player"
	nbateam "github.com/Sirius198/go-nba-sports-api-test/nba/team"
)

func main() {

	loadEnv()
	executeCronJobs()
	nbadatabase.ConnectDB()

	r := gin.Default()
	r.GET("/GamesByDate/:date", nbagame.GetGamesByDate)
	r.GET("/PlayersByTeam/:team", nbaplayer.GetPlayersByTeam)
	r.GET("/PlayerDetailsByName/:firstname/:lastname", nbaplayer.GetPlayerDetailsByName)

	r.Run()
	// listen and serve on 0.0.0.0:8080
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func executeCronJobs() {
	gocron.Every(4).Hours().Do(nbateam.FetchAllTeams)
	gocron.Every(4).Hours().Do(nbaplayer.FetchAllPlayers)
	gocron.Start()
}
