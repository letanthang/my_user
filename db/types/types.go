package types

import "time"

type LoginReq struct {
	Phone    string `json:"phone" bson:"phone"`
	Password string `json:"password" bson:"password"`
}

type AccountSearchReq struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ClassName string `json:"class_name"`
}

type AccountAddReq struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
	Password  string `json:"password" bson:"password"`
}

type AccountUpdateReq struct {
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	Email     string `json:"email,omitempty" bson:"email"`
	Password  string `json:"password,omitempty" bson:"password"`
}

type Account struct {
	ID        int        `json:"id" bson:"id"`
	FirstName string     `json:"first_name" bson:"first_name"`
	LastName  string     `json:"last_name" bson:"last_name"`
	Email     string     `json:"email" bson:"email"`
	Phone     string     `json:"phone" bson:"phone"`
	Password  string     `json:"password" bson:"password" groups:"private"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
}

type DeleteReq struct {
	ID int `json:"id"`
}
