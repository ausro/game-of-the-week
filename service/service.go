package service

import (
	"context"
	"errors"
	"log"

	"github.com/ausro/game-of-the-week/db"
)

type SteamAppService struct {
	db  db.Database
	env string
}

func NewSteamAppService(db db.Database, env string) *SteamAppService {
	return &SteamAppService{db, env}
}

func (steamAppServ *SteamAppService) AddSteamApp(ctx context.Context, steamApp *db.SteamAppModel) error {
	_, err := steamAppServ.db.Insert(ctx, steamApp)
	if err != nil {
		log.Printf("Failed to Insert: %s", err)
		return err
	}

	return nil
}

func (steamAppServ *SteamAppService) DeleteSteamApp(ctx context.Context, steamApp *db.SteamAppModel) error {
	_, err := steamAppServ.db.Delete(ctx, "app", db.Filter{"id": steamApp.ID})
	if err != nil {
		log.Printf("Failed to Delete: %s", err)
		return err
	}

	return nil
}

func (steamAppServ *SteamAppService) GetAllSteamApps(ctx context.Context) (*[]db.SteamAppModel, error) {
	var steamAppList []db.SteamAppModel
	err := steamAppServ.db.SelectAll(ctx, &steamAppList, "app")
	if err != nil {
		log.Printf("Failed to select from DB: %s", err)
		return nil, err
	}

	if len(steamAppList) == 0 {
		log.Printf("App list is empty.")
		return nil, errors.New("db_no_rows")
	}

	return &steamAppList, err
}

func (steamAppServ *SteamAppService) GetPromotedApps(ctx context.Context) (*[]db.SteamAppModel, error) {
	var promotedList []db.SteamAppModel
	err := steamAppServ.db.Select(ctx, &promotedList, "promoted", true)
	if err != nil {
		log.Printf("Failed to select promoted from DB: %s", err)
	}

	return &promotedList, err
}

func (steamAppServ *SteamAppService) ToggleAppPromoted(ctx context.Context, steamApp *db.SteamAppModel) error {
	steamApp.Promoted = !steamApp.Promoted

	err := steamAppServ.AddSteamApp(ctx, steamApp)
	if err != nil {
		return err
	}

	return nil
}
