package Services


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	// Dsn: "user:password@tcp(127.0.0.1:3306)/test"
	dsn := "root:S12p_w99Q@tcp(139.198.5.192:3306)/cargocvge_auth"
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