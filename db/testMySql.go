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
	//dsn := "root:S12p_w99Q@tcp(139.198.5.192:3306)/test"
	dsn := "root:S12p_w99Q@tcp(139.198.5.192:3306)/investorsdata"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}

	db.SetMaxOpenConns(10)

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

type Student struct {
	sno int
	sname string
	ssex string
	sage int
	sdept string
}

func GetData(i int) Student {
	dbBase := GetDB()

	var t Student

	s := `SELECT * FROM student WHERE Sno = ?`
	err = dbBase.QueryRow(s, i).Scan(&t.sno, &t.sname, &t.ssex, &t.sage, &t.sdept)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dbBase.Stats())

	return t
}
