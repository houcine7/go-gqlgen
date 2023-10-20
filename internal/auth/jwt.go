package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	SecretKey = []byte("Hello world _ Secret")
)

// To generate a jwt token using username as the object
func GenerateToken(username string) (string,error) {
	token := jwt.New(jwt.SigningMethodHS256) // create token with the signing method
	claims := token.Claims.(jwt.MapClaims) // Map to store our claims

	//set claims
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// token 
	tokenStr,err := token.SignedString(SecretKey)

	if err != nil{
		log.Fatal("Error in key generation",err)
		return "",err
	}

	return tokenStr, nil
}

// parse the jwt
func ParseToken(tokenStr string) (string,error) {

	// anonymous func
	keyFunc := func(token *jwt.Token) (interface{},error){
		return SecretKey,nil
	}

	token, err := jwt.Parse(tokenStr,keyFunc)

	if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username :=claims["username"].(string)
		return username,nil
	}
	return "",err
}



