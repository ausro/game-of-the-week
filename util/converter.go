package util

import (
	"log"

	"github.com/ausro/game-of-the-week/api"
	"github.com/ausro/game-of-the-week/db"
)

func DetailToApp(detail *api.AppDetails) *db.SteamAppModel {
	var app db.SteamAppModel
	var genres []string
	var screenshots []string

	if detail.IsFree {
		app.Price = "Free"
	} else {
		app.Price = detail.PriceOverview.FinalFormatted
	}

	app.ID = detail.SteamAppId
	app.Name = detail.Name
	app.ReleaseDate = detail.ReleaseDate.Date
	app.ShortDescription = detail.ShortDescription
	app.HeaderImage = detail.HeaderImage
	app.Promoted = false

	for i := range detail.Genres {
		log.Printf("Adding genre: %s for %s", detail.Genres[i].Description, detail.Name)
		genres = append(genres, detail.Genres[i].Description)
	}

	for j := range detail.Screenshots {
		screenshots = append(screenshots, detail.Screenshots[j].PathThumbnail)
	}

	app.Genres = genres
	app.Screenshots = screenshots

	return &app
}
