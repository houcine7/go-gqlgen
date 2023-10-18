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

	for _, env := range os.Environ() {
		fmt.Println(env)
	}	

	connectionStr := fmt.Sprintf("root:%s@tcp(localhost)/go_graphql",os.Getenv("MYSQL_DUMMY_PASSWD"))
	log.Print(os.Getenv("MYSQL_DUMMY_PASSWD"))
	db, err := sql.Open("mysql", connectionStr)

	if err !=nil {
		log.Panic(err)
	}

	if err=db.Ping() ; err !=nil{
		log.Panic(err)
	}
	//log.Fatal(db.Stats())
	Db = db;
}

func CloseDB() error{
	return Db.Close();
}
