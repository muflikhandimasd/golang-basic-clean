package configs

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDBSQL(c SQLConfig) *sql.DB {

	db, err := sql.Open("mysql", fmt.Sprint(c.DB_USER, ":", c.DB_PASS, "@tcp(", c.DB_HOST, ":", c.DB_PORT, ")/", c.DB_NAME, "?parseTime=True&loc=Asia%2FJakarta&charset=utf8&autocommit=false"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("DATABASE CONNECTED " + c.DB_NAME)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
