package user

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	passwordHash string `json:"-" bson:"password"`
	Email        string `json:"email" bson:"email"`
}

type CreateUserDTO struct {
	username string `json:"username"`
	password string `json:"passsword"`
	Email    string `json:"email"`
}
