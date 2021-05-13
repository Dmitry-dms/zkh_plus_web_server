package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	Id           primitive.ObjectID `json:"-" db:"company_id" bson:"_id"`
	Email        string             `json:"email" binding:"required" db:"email" bson:"email"`
	Password     string             `json:"password" binding:"required" db:"password_hash" bson:"password"`
	Name         string             `json:"name" binding:"required" db:"company_name" bson:"name"`
	DirectorName string             `json:"director_full_name" binding:"required" db:"director_full_name" bson:"director_full_name"`
	Phone        string             `json:"company_phone" binding:"required" db:"company_phone" bson:"company_phone"`
	City         string             `json:"company_city" binding:"required" db:"company_city" bson:"company_city"`
	Street       string             `json:"company_street" binding:"required" db:"company_street" bson:"company_street"`
	HomeNumber   string             `json:"company_home_number" binding:"required" db:"company_home_number" bson:"company_home_number"`
	Flat         string             `json:"company_flat" binding:"required" db:"company_flat" bson:"company_flat"`
}
