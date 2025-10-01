package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"github.com/arun14k08/goframework/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB 
var DBConnector *Queries 

func Connect() error {
    dbConn, err := sql.Open("postgres", config.AppProp.DBUrl)
    if err != nil {
        return fmt.Errorf("failed to open database: %w", err)
    }

    // Test the connection
    if err := dbConn.PingContext(context.Background()); err != nil {
        return fmt.Errorf("cannot ping database: %w", err)
    }

    DB = dbConn
    DBConnector = New(DB)

    log.Println("Database connected successfully")
    return nil
}

func Close() {
    if DB != nil {
        DB.Close()
    }
}
