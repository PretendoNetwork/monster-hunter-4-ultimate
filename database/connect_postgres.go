package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/PretendoNetwork/monster-hunter-4-ultimate/globals"
)

var Postgres *sql.DB

func ConnectPostgres() {
	var err error

	Postgres, err = sql.Open("postgres", os.Getenv("PN_MH4U_POSTGRES_URI"))
	if err != nil {
		globals.Logger.Critical(err.Error())
	}

	globals.Logger.Success("Connected to Postgres!")
}
