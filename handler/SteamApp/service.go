package steamapp

import (
	"context"
	"time"

	"github.com/ausro/game-of-the-week/db"
	"github.com/ausro/game-of-the-week/handler"
	"github.com/ausro/game-of-the-week/service"
	"github.com/gin-gonic/gin"
)

type SteamAppHandler struct {
	Server  *handler.Server
	group   string
	router  *gin.RouterGroup
	Service *service.SteamAppService
	Apps    *AppList
}

type AppList struct {
	appList         map[int]db.SteamAppModel
	recommendedList []db.SteamAppModel
	expiry          time.Time
}

func NewSteamAppHandler(server *handler.Server, groupName string, service *service.SteamAppService) *SteamAppHandler {
	var apps = AppList{
		appList:         map[int]db.SteamAppModel{},
		recommendedList: []db.SteamAppModel{},
		expiry:          time.Now().AddDate(0, 0, -1),
	}

	steamAppHandler := &SteamAppHandler{
		server,
		groupName,
		&gin.RouterGroup{},
		service,
		&apps,
	}

	steamAppHandler.router = steamAppHandler.registerGroup()
	steamAppHandler.routes()
	steamAppHandler.init(context.Background())

	return steamAppHandler
}
