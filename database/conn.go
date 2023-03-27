package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Connect to database
func Connect() *sql.DB {

    env := func(key string) string {
        return os.Getenv(key)
    }

    //user file .env
    db, err := sql.Open("mysql", env("DB_USER")+":"+env("DB_PASSWORD")+"@tcp("+env("DB_HOST")+":"+env("DB_PORT")+")/"+env("DB_DATABASE"))
    if err != nil {
        fmt.Println("Error connecting to database")
    }
    fmt.Println("Connected to database!")

    return db
}

// Close connection
func Close(db *sql.DB) {
    db.Close()
}
