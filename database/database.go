package database

import (
	"database/sql"
	"fmt"
	c "go-rest-api/config"
	"go-rest-api/helpers"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {

	config := c.GetDbConfig()
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	var err error
	db, err = sql.Open("mysql", dbSource)

	helpers.Catch(err)
}

// GetDb ...
func GetDb() *sql.DB {

	initDB()

	return db
}
