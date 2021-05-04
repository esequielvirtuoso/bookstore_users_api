package users_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlURL = "MYSQL_URL"  // root:admin123@tcp(127.0.0.1:3306)/users_db?charset=utf8
)

var (
	Client *sql.DB
)

func init() {
	var err error
	Client, err = sql.Open("mysql", getMySQLURL())
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}

func getMySQLURL() string {
	return os.Getenv(mysqlURL)
}