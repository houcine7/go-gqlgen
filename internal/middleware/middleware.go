package auth

import (
	"context"
	"net/http"
	"strconv"

	jwt "github.com/houcine7/graphql-server/internal/auth"
	"github.com/houcine7/graphql-server/internal/models/users"
)


type contextKey struct{
	name string
}


var userCtxKey = &contextKey{"user"}


func AuthMiddleware() func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resWriter http.ResponseWriter, req *http.Request) {
			//
			header := req.Header.Get("Authorization")
			if header =="" { //unauthenticated users 
				next.ServeHTTP(resWriter,req)
				return 
			}
			// 
			tokenStr := header
			username, err :=jwt.ParseToken(tokenStr)
			if err!=nil{
				http.Error(resWriter,"You are not allowed",http.StatusForbidden)
				return
			}
			user := users.User{Username: username}
			id,err := users.GetUserId(username)

			if err!=nil{
				next.ServeHTTP(resWriter,req)
				return
			}
			user.ID = strconv.Itoa(id)
			ctx := context.WithValue(req.Context(), userCtxKey,&user) // add user to context 
			
			req = req.WithContext(ctx)
			next.ServeHTTP(resWriter,req)
		})
	}
}
//below function is used to get the user from the context
// requires a middleware to be used before
func ForContext(ctx context.Context) *users.User{
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}