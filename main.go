package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/matheuspolitano/GoMasterDB/api"
	db "github.com/matheuspolitano/GoMasterDB/db/sqlc"
)

const (
	dbDriver  = "postgres"
	conString = "postgresql://root:MP@TEST123@localhost:5432/fintech_db?sslmode=disable"
)

func main() {
	con, err := sql.Open(dbDriver, conString)
	if err != nil {
		log.Fatal(err)
	}
	query := db.New(con)

	server := api.NewServer(query)

	if err = server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
