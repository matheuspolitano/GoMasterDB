package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/matheuspolitano/GoMasterDB/utils"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conf, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(conf.DbDriver, conf.DbConnetionString)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
