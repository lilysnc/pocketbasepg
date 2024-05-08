//go:build !sqlite3

package core

import (
	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
)

func connectDB(dsn string) (*dbx.DB, error) {
	return dbx.MustOpen("postgres", dsn)
}
