package main

import (
	"log"

	"github.com/Sirius198/go-nba-sports-api-test/nba"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadEnv()

	r := gin.Default()
	r.GET("/GamesByDate/:date", nba.GetGamesByDate)
	r.GET("/PlayersByTeam/:team", nba.GetPlayersByTeam)
	r.GET("/PlayerDetailsByName/:firstname/:lastname", nba.GetPlayerDetailsByName)

	r.Run()
	// listen and serve on 0.0.0.0:8080
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
