package users

import (
	"database/sql"
	"fmt"
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
	stm, err := database.Db.Prepare("INSERT INTO User(Username, Password) VALUES(?,?)")
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

func (user *User) Authenticate() (bool, error){
	qr,err := database.Db.Prepare("SELECT Password FROM User WHERE Username=?")
	
	if err!=nil{
		fmt.Print("Error occurred",err)
		return false, fmt.Errorf("unexpected error ! :)")
	}
	defer qr.Close()

	row := qr.QueryRow(user.Username)
	var dbPasswd string
	err = row.Scan(&dbPasswd)

	if err!=nil{
		fmt.Print("Can't find provided username",err)
		if err ==sql.ErrNoRows{
			return false,fmt.Errorf("username not found")
		}else{
			return false, fmt.Errorf("unexpected error ! :)")
		}
	}

	if CheckPasswordHash(user.Password,dbPasswd) {
		return true, nil
	}else{
		return false, fmt.Errorf("wrong password try again")
	}

}


func GetUserId(username string) (int,error) {
	stmt, err := database.Db.Prepare("SELECT ID FROM User WHERE Username=?")

	if err !=nil{
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
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



//utils : =>

func HashPassword(password string) (string,error){
	bytes,err :=bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}