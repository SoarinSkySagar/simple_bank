package main

import (
	"database/sql"
	"log"

	"github.com/SoarinSkySagar/simple_bank/api"
	db "github.com/SoarinSkySagar/simple_bank/db/sqlc"
	"github.com/SoarinSkySagar/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load env variables:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
