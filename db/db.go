package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Database interface {
	Insert(ctx context.Context, model any) (sql.Result, error)
	Select(ctx context.Context, model any, column string, param any) error
	SelectAll(ctx context.Context, model any, table string) error
	SelectWithMultipleFilter(ctx context.Context, model any, condition Filter) error
	Delete(ctx context.Context, tableName string, filter Filter) (sql.Result, error)
	Update(ctx context.Context, tableName string, Set Filter, Condition Filter) (sql.Result, error)
	Raw(ctx context.Context, model any, query string, args ...any) error
	Migrate() error
	Close() error
}

type Filter map[string]any

type DB struct {
	db *bun.DB
}

func (d *DB) whereCondition(filter Filter, ConditionType string) string {
	var whereClauses []string
	for key, val := range filter {
		var formattedVal string
		switch v := val.(type) {
		case string:
			formattedVal = fmt.Sprintf("'%s'", v)
		case int, int64:
			formattedVal = fmt.Sprintf("%d", v)
		case float64:
			formattedVal = fmt.Sprintf("%.2f", v)
		default:
			log.Fatal("Unhandled WHERE condition type.")
		}

		whereClauses = append(whereClauses, fmt.Sprintf("%s=%s", key, formattedVal))
	}

	var result string
	if len(whereClauses) > 0 {
		switch ConditionType {
		case "SET":
			result = strings.Join(whereClauses, " , ")
		case "AND":
			result = strings.Join(whereClauses, " AND ")
		}
	}

	return result
}

func (d *DB) Delete(ctx context.Context, tableName string, filter Filter) (sql.Result, error) {
	return d.db.NewDelete().Table(tableName).Where(d.whereCondition(filter, "AND")).Exec(ctx)
}

func (d *DB) Insert(ctx context.Context, model any) (sql.Result, error) {
	return d.db.NewInsert().Model(model).On("CONFLICT (id) DO UPDATE").Exec(ctx)
}

func (d *DB) Migrate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	models := []any{
		(*SteamAppModel)(nil),
	}

	for _, model := range models {
		if _, err := d.db.NewCreateTable().Model(model).WithForeignKeys().IfNotExists().Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (d *DB) Select(ctx context.Context, model any, column string, param any) error {
	return d.db.NewSelect().Model(model).Where(fmt.Sprintf("%s = ?", column), param).Scan(ctx)
}

func (d *DB) SelectAll(ctx context.Context, model any, table string) error {
	return d.db.NewSelect().Table(table).Scan(ctx, model)
}

func (d *DB) SelectWithMultipleFilter(ctx context.Context, model any, condition Filter) error {
	return d.db.NewSelect().Model(model).Where(d.whereCondition(condition, "AND")).Scan(ctx)
}

func (d *DB) Update(ctx context.Context, tableName string, Set Filter, Condition Filter) (sql.Result, error) {
	return d.db.NewUpdate().Table(tableName).Set(d.whereCondition(Set, "SET")).Where(d.whereCondition(Condition, "AND")).Exec(ctx)
}

func (d *DB) Raw(ctx context.Context, model any, query string, args ...any) error {
	return d.db.NewRaw(query, args...).Scan(ctx, model)
}

func (d *DB) Close() error {
	slog.Info("Closing DB connection.")
	return d.db.Close()
}

func New() Database {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Invalid database port")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUsername, dbPassword, dbHost, port, dbName)
	database := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))

	db := bun.NewDB(database, pgdialect.New())
	return &DB{db: db}
}
