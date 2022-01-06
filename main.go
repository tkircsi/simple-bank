package main

import (
	"database/sql"
	"log"

	"github.com/tkircsi/simple-bank/api"
	db "github.com/tkircsi/simple-bank/db/sqlc"
	"github.com/tkircsi/simple-bank/util"

	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read configuration:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start HTTP server:", err)
	}
}
