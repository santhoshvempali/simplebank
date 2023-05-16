package main

import (
	"database/sql"
	"log"

	"github.com/santhoshvempali/simplebank/api"
	db "github.com/santhoshvempali/simplebank/db/sqlc"
	"github.com/santhoshvempali/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read from file", err)
	}
	conn, err := sql.Open(config.DB_DRIVER, config.DB_SERVICE)
	if err != nil {
		log.Fatal("error connection to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.SERVICE_ADDRESS)
	if err != nil {
		log.Fatal("Cannot start the server", err)
	}

}
