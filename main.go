package main

import (
	"database/sql"
	"log"

	"github.com/go_dev/simplebank/api"
	db "github.com/go_dev/simplebank/db/sqlc"
	"github.com/go_dev/simplebank/utils"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDrvier, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
