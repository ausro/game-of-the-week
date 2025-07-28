package steamapp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ausro/game-of-the-week/api"
	"github.com/ausro/game-of-the-week/db"
	"github.com/ausro/game-of-the-week/util"
	"github.com/gin-gonic/gin"
)

func (h *SteamAppHandler) registerGroup() *gin.RouterGroup {
	return h.Server.Gin.Group(h.group)
}

func (h *SteamAppHandler) init(ctx context.Context) {
	h.ListApps(ctx)
}

func (h *SteamAppHandler) routes() http.Handler {
	h.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, h.ListApps(c))
	})
	h.router.GET("/recommended", func(c *gin.Context) {
		c.JSON(http.StatusOK, h.ListRecommended(c))
	})

	return h.Server.Gin
}

func (h *SteamAppHandler) Add(ctx context.Context, app *db.SteamAppModel) error {
	err := h.Service.AddSteamApp(ctx, app)
	if err != nil {
		log.Printf("Failed to add steam app: %s, %s", app.Name, err)
		return err
	}

	h.Apps.appList[app.ID] = *app
	if app.Promoted {
		h.Apps.recommendedList = append(h.Apps.recommendedList, *app)
	}

	return nil
}

func (h *SteamAppHandler) GetAppById(id int) *db.SteamAppModel {
	app := h.Apps.appList[id]
	if app.ID == 0 {
		return nil
	}

	return &app
}

func (h *SteamAppHandler) delete(ctx context.Context, app *db.SteamAppModel) error {
	err := h.Service.DeleteSteamApp(ctx, app)
	if err != nil {
		log.Printf("Failed to delete steam app: %d, %s", app.ID, err)
		return err
	}

	delete(h.Apps.appList, app.ID)

	return nil
}

func (h *SteamAppHandler) ListApps(ctx context.Context) map[int]db.SteamAppModel {
	if len(h.Apps.appList) == 0 {
		appList, err := h.Service.GetAllSteamApps(ctx)
		if err != nil {
			if appList == nil {
				h.CreateDefaultRecommended(ctx)
				return h.Apps.appList
			} else {
				return nil
			}
		}

		if appList == nil || len(*appList) == 0 {
			log.Fatalf("No games exist and unable to pull games, exiting: %s", err)
		}

		for _, app := range *appList {
			h.Apps.appList[app.ID] = app

			if app.Promoted {
				h.Apps.recommendedList = append(h.Apps.recommendedList, app)
			}
		}
	}

	if h.Apps.expiry.Before(time.Now()) {
		err := h.RequestGames(ctx)
		if err != nil {
			log.Printf("Unable to request games: %s", err)
		}

		h.Apps.expiry = time.Now().AddDate(0, 0, 1)
	}

	return h.Apps.appList
}

func (h *SteamAppHandler) GetBlacklist(ctx context.Context) map[int]db.BlacklistModel {
	if len(h.Blacklist.ids) == 0 {
		blacklist, err := h.Service.GetBlacklist(ctx)
		if err != nil {
			return nil
		}

		for _, app := range *blacklist {
			h.Blacklist.ids[app.ID] = app
		}
	}

	return h.Blacklist.ids
}

func (h *SteamAppHandler) CreateDefaultRecommended(ctx context.Context) {
	reccIds := [...]int{3527290, 3241660, 1966720, 739630}

	log.Printf("Creating default apps: %d", reccIds)

	for _, id := range reccIds {
		go h.AddGameById(ctx, id)
	}
}

func (h *SteamAppHandler) ListRecommended(ctx context.Context) *[]db.SteamAppModel {
	if len(h.Apps.recommendedList) != 0 {
		return &h.Apps.recommendedList
	} else {
		appList, err := h.Service.GetPromotedApps(ctx)
		if err != nil {
			log.Printf("Failed to list promoted apps: %s", err)
			return nil
		}

		return appList
	}
}

func (h *SteamAppHandler) RequestGames(ctx context.Context) error {
	resp, err := api.GET(api.GetApps, "category=cat_newreleases&cc=us&l=english")
	if err != nil {
		log.Printf("failed to GET: %s", err)
		return err
	}

	appsCat, err := util.GetAsAppsCategory(resp)
	if err != nil {
		return err
	}

	cat := appsCat.Tabs["viewall"]

	for _, item := range cat.Items {
		go h.AddGameById(ctx, item.Id)
	}

	return nil
}

func (h *SteamAppHandler) AddGameById(ctx context.Context, id int) {
	if h.Blacklist.ids[id].ID != 0 {
		return
	}

	dResp, err := api.GET(api.GetDetails, fmt.Sprintf("appids=%d", id))
	if err != nil {
		log.Printf("error getting details, skipping: %s", err)
		return
	}

	detail, err := util.GetAsAppDetails(dResp)
	if err != nil {
		log.Printf("Failed to parse app details: %s", err)
		return
	}

	err = util.ValidateApp(detail)
	if err != nil {
		log.Printf("Skipping %s: %s", detail.Name, err)
		return
	}

	app := util.DetailToApp(detail)
	h.Add(ctx, app)
}

func (h *SteamAppHandler) DeleteGameById(ctx context.Context, id int) {
	err := h.delete(ctx, &db.SteamAppModel{ID: id})
	if err != nil {
		log.Printf("Game of ID %d could not be removed, verify the id is correct.", id)
	}
}

func (h *SteamAppHandler) BlacklistGameById(ctx context.Context, id int) {
	appId := &db.BlacklistModel{ID: id, Name: fmt.Sprint(id)}

	err := h.Service.BlacklistSteamApp(ctx, appId)
	if err != nil {
		log.Printf("Game of ID %d could not be blacklisted due to %s", id, err)
	}

	h.Blacklist.ids[id] = *appId
}
