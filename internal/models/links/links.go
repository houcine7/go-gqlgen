package links

import (
	"database/sql"
	"log"

	database "github.com/houcine7/graphql-server/internal/db"
	"github.com/houcine7/graphql-server/internal/models/users"
)

type Link struct {
	ID      string `json:"id"`
	Title   string
	Address string
	User    *users.User
}

func (l Link) Save() int64{
	qr , err := database.Db.Prepare("INSERT INTO Link(Title,address) values(?,?)")

	if err!=nil{
		log.Fatal(err)
	}
	defer qr.Close()
	res,err := qr.Exec(l.Title,l.Address)
	if err!=nil{
		log.Fatal(err)
	}
	
	id,err :=res.LastInsertId()

	if err!=nil{
		log.Fatal("Error ...: ",err.Error())
	}
	log.Print("Row Inserted successfully",id)
	return id
}

func Links() ([]Link,error ){

	myQuery,err := database.Db.Prepare("SELECT * FROM Link");
	if err !=nil {
		log.Fatal(err)
	}

	defer myQuery.Close() // schedule this to be executed at last (first deferred to be executed )

	rows, err := myQuery.Query()
	if err !=nil{
		log.Fatal(err)
	}
	
	defer rows.Close() // second deferred to execute 
	var links []Link 

	for rows.Next() {
		var title,address string
		var id int64
		var userId sql.NullInt64
		
		if err:= rows.Scan(&id,&title,&address,&userId); err!=nil{
			log.Fatal(err)
		}

		link := Link{
			Title: title,
			Address: address,
			ID: string(id),
		}
		links =append(links,link)
	}
	
	return links,nil

}