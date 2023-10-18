package links

import (
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