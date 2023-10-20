package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)


var Db *sql.DB

func InitDb() {	

	connectionStr := fmt.Sprintf("root:%s@tcp(localhost)/go_graphql",os.Getenv("MYSQL_DUMMY_PASSWD"))
	log.Println(connectionStr)
	db, err := sql.Open("mysql", connectionStr)

	if err !=nil {
		log.Panic(err)
	}

	if err=db.Ping() ; err !=nil{
		log.Panic(err)
	}
	Db = db;
}

func CloseDB() error{
	return Db.Close();
}
