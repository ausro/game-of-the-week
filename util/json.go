package util

import (
	"encoding/json"

	"github.com/ausro/game-of-the-week/api"
)

func GetAsAppDetails(js []byte) (*api.AppDetails, error) {
	var resp api.SteamAppResponse
	var details api.AppDetails

	err := json.Unmarshal(js, &resp)

	for i := range resp {
		details = resp[i].Data
	}

	return &details, err
}

func GetAsAppsCategory(js []byte) (api.AppsCategory, error) {
	var category api.AppsCategory

	err := json.Unmarshal(js, &category)

	return category, err
}
