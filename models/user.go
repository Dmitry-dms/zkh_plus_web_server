package models



import "go.mongodb.org/mongo-driver/bson/primitive"

//Аккаунт
type User struct {
	Id         primitive.ObjectID   `json:"-" db:"user_id" bson:"_id"`
	Name       string `json:"name" binding:"required" bson:"name"`
	Surname    string `json:"surname" binding:"required" bson:"surname"`
	Patronymic string `json:"patronymic" binding:"required" bson:"patronymic"`
	FullName   string	`bson:"full_name"`
	Email      string `json:"email" binding:"required" bson:"email"`
	Password   string `json:"password" binding:"required" bson:"password"`
	CompanyId  primitive.ObjectID    `json:"-" db:"company_id" bson:"company_id"`
}

