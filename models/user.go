package models

//Аккаунт
type User struct {
	//gorm.Model
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	//Token    string `json:"token" ;sql:"-"`
}
