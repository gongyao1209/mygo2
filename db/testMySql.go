package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	// Dsn: "user:password@tcp(127.0.0.1:3306)/test"
	dsn := "gongyao:Passw0rd@tcp(127.0.0.1:3306)/learning"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(1)

	db.Ping()
}

func GetDB() *sql.DB {
	return db
}

type Temp struct {
	Id      int    `json:"id"`
	Student string `json:"student"`
	Class   string `json:"class"`
}

func GetData() Temp {
	var t Temp

	var id int
	var student, class string

	s := `SELECT * FROM courses WHERE id = 10`
	dbBase := GetDB()
	dbBase.QueryRow(s).Scan(&id, &student, &class)

	t.Id = id
	t.Student = student
	t.Class = class

	fmt.Println(t)

	return t
}
