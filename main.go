package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sim_bank/api"
	dbs "sim_bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "127.0.0.1:8081"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot connect to dba")
	}
	store := dbs.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalln("Error while starting the server", err)
	}

}
