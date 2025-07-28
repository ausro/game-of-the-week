package db

import (
	"github.com/uptrace/bun"
)

type BlacklistModel struct {
	bun.BaseModel `bun:"table:blacklist"`

	ID   int    `bun:"id,pk" json:"id"`
	Name string `bun:",notnull" json:"name"`
}
