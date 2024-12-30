package main

import (
	"database/sql"
	"log"

	"github.com/fuad7161/simple-bank/api"
	db "github.com/fuad7161/simple-bank/db/sqlc"
	"github.com/fuad7161/simple-bank/db/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("connot start server: ", err)
	}
}
