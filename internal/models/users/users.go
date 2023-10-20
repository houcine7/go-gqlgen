package users

import (
	"database/sql"
	"log"

	database "github.com/houcine7/graphql-server/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"` // when you encode/decode with GO JSON use this field name
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Create() int64{
	stm, err := database.Db.Prepare("INSERT INTO Users(Username, Password) VALUES(?,?)")
	if err!=nil{
		log.Fatal(err)
	}
	defer stm.Close()

	hashedPassword ,err:= HashPassword(user.Password)
	if err!=nil{
		log.Fatal(err)
	}

	res ,err := stm.Exec(user.Username,hashedPassword)

	if err!=nil{
		log.Fatal(err)
	}
	id,err := res.LastInsertId()
	
	if err!=nil{
		log.Fatal(err)
	}
	return id
}

func HashPassword(password string) (string,error){
	bytes,err :=bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}

func GetUserId(username string) (int,error) {
	stmt, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username=?")

	if err !=nil{
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	if err !=nil{
		if err !=sql.ErrNoRows{
			log.Print(err)
		}
		return 0,err	
	}

	return id,nil

}	