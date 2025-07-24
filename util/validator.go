package util

import (
	"errors"

	"github.com/ausro/game-of-the-week/api"
)

// Tag(s) of: Online Co-op (extra credit to action and horror).
// Price: <$30.
// Released or early access (not coming soon).
func ValidateApp(appDetail *api.AppDetails) error {
	if appDetail.SteamAppId == 0 {
		return errors.New("invalid steamapp id")
	}

	if appDetail.PriceOverview.Initial > 3000 {
		return errors.New("price above limit")
	}

	if appDetail.ReleaseDate.ComingSoon {
		return errors.New("game not released")
	}

	for _, cat := range appDetail.Categories {
		if cat.Description == "Online Co-op" {
			return nil
		}
	}

	return errors.New("game does not meet criteria")
}
