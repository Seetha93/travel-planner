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

func Insert(db *sql.DB, tableName string, name string, code string) {
	id := 0
	sqlStatement := `
			INSERT INTO ` + tableName + ` (name, code, created_at)
			VALUES ($1, $2, $3)
			RETURNING id`
	err := db.QueryRow(sqlStatement, name, code, time.Now()).Scan(&id)

	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func getId(db *sql.DB, sqlStatement string) (id int) {
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return -1
	case nil:
		fmt.Println(id)
		return id
	default:
		panic(err)
	}
}

func InsertConnectingData(db *sql.DB, tableName string,
	parent string, child string,
	parentTableName string, childTableName string,
	parentCoulmnName string, childColumnName string) {
	id := 0
	sqlStatement := `select id from ` + childTableName + ` where code = '` + child + `';`
	childId := getId(db, sqlStatement)

	sqlStatement = `select id from ` + parentTableName + ` where code = '` + parent + `';`
	parentId := getId(db, sqlStatement)
	sqlStatement = `
			INSERT INTO ` + tableName + ` (` + parentCoulmnName + `, ` + childColumnName + `, created_at)
			VALUES ($1, $2, $3)
			RETURNING id`
	err := db.QueryRow(sqlStatement, parentId, childId, time.Now()).Scan(&id)

	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func CloseConnection(db *sql.DB) {
	db.Close()
	fmt.Println("Closed the DB connection")
}
