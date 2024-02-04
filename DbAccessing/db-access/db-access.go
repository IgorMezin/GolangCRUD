package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/alexbrainman/odbc"
)

var db *sql.DB
var err error
var conn *odbc.Conn

type ConnectionString struct {
	server   string
	port     int
	user     string
	password string
	database string
	driver   string
}

var cs ConnectionString

func main() {
	// Capture connection properties
	cs := ConnectionString{server: "127.0.0.1", port: 1433, user: "sa", password: "1234", database: "master", driver: "{ODBC Driver 17 for SQL Server}"}
	connString := fmt.Sprintf("server=%s;uid=%s;pwd=%s;port=%d;database=%s;driver=%s",
		cs.server,
		cs.user,
		cs.password,
		cs.port,
		cs.database,
		cs.driver)
	// Connection Begin
	db, err = sql.Open("odbc", connString)
	if err != nil {
		log.Fatal("Error connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	serverVersion()

	//Create test table
	//createId, err = CreateTable("test", "")

}

func serverVersion() {
	// Use background context
	ctx := context.Background()

	// ping database to see if it's still alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}

	var result string

	err = db.QueryRowContext(ctx, "Select @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed: ", err.Error())
	}
	log.Printf("%s\n", result)
}
