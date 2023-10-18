package users

type User struct{
	ID string `json:"id"` // when you encode/decode with GO JSON use this field name
	Username string `json:"username"`
	Password string `json:"password"`
}