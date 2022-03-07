package persistence

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewDBConnection(url string) (*sql.DB, error) {
	return sql.Open("pgx", url)
}
