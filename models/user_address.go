package models

type UserAddress struct {
	Id         int    `json:"-" db:"address_id"`
	UserId     int    `json:"-" db:"user_id"`
	City       string `json:"city" binding:"required" db:"city"`
	Street     string `json:"street" binding:"required" db:"street"`
	HomeNumber string `json:"home_number" binding:"required" db:"home_number"`
	Flat       string `json:"flat" binding:"required" db:"flat"`
}
