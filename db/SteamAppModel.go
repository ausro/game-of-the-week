package db

import (
	"github.com/uptrace/bun"
)

type SteamAppModel struct {
	bun.BaseModel `bun:"table:app"`

	ID               int      `bun:"id,pk" json:"id"`
	Name             string   `bun:",notnull" json:"name"`
	Price            string   `bun:",notnull" json:"price"`
	Genres           []string `bun:",array" json:"genres"`
	ReleaseDate      string   `bun:",notnull" json:"release_date"`
	ShortDescription string   `bun:",notnull" json:"short_description"`
	HeaderImage      string   `bun:",notnull" json:"header_image"`
	Screenshots      []string `bun:",notnull" json:"screenshots"`
	Promoted         bool     `bun:",notnull" json:"promoted"`
}
