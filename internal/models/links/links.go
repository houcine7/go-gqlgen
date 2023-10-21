package links

import (
	"log"
	"strconv"

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
	qr , err := database.Db.Prepare("INSERT INTO Link(Title,Address,UserID) values(?,?,?)")

	if err!=nil{
		log.Fatal(err)
	}
	defer qr.Close()
	res,err := qr.Exec(l.Title,l.Address,l.User.ID)
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

	myQuery,err := database.Db.Prepare("SELECT l.ID,l.Title,l.Address,U.ID,U.Username FROM Link as l INNER JOIN User as U on l.UserID = U.ID");
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
		var link Link
		var username string
		var id int
		
		if err:= rows.Scan(&link.ID,&link.Title,&link.Address,&id,&username); err!=nil{
			log.Fatal(err)
		}

		link.User = &users.User{
			ID: strconv.Itoa(id),
			Username: username,
		}
		links =append(links,link)
	}
	
	return links,nil
}