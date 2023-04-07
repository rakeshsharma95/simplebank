package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rdevelop/simplebank/api"
	db "github.com/rdevelop/simplebank/db/sqlc"
	"github.com/rdevelop/simplebank/util"
)

// const (
//
//	dbDriver      = "postgres"
//	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
//	serverAddress = "0.0.0.0:8080"
//
// )
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.DBAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
