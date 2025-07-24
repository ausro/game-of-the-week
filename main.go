package main

import (
	"log"
	"os"

	"github.com/ausro/game-of-the-week/db"
	"github.com/ausro/game-of-the-week/handler"
	command "github.com/ausro/game-of-the-week/handler/Command"
	steamapp "github.com/ausro/game-of-the-week/handler/SteamApp"
	"github.com/ausro/game-of-the-week/service"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	env := os.Getenv("APP_ENV")

	db := db.New()
	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	serv := handler.New(db)

	steamAppService := service.NewSteamAppService(db, env)
	handler := steamapp.NewSteamAppHandler(serv, "/", steamAppService)

	command.NewCommandHandler(handler)

	log.Fatal(serv.Run())
}
