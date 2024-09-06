package main

import (
	"database/sql"
	"log"
	"simplebank-go/api"
	db "simplebank-go/db/sqlc"
	"simplebank-go/util"

	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBdriver, config.DBsource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
