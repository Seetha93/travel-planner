package database

import (
	"database/sql"
	"fmt"
	"time"

	Config "travel-planner/utils"

	_ "github.com/lib/pq"
)

func Connect() (db *sql.DB, err error) {
	config := Config.LoadDatabaseConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected")
	return db, err
}

func Insert(db *sql.DB, name string, code string) {
	id := 0
	sqlStatement := `
			INSERT INTO country (name, code, created_at)
			VALUES ($1, $2, $3)
			RETURNING id`
	err := db.QueryRow(sqlStatement, name, code, time.Now()).Scan(&id)

	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func CloseConnection(db *sql.DB) {
	db.Close()
	fmt.Println("Closed the DB connection")
}
