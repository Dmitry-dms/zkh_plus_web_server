package models

//Аккаунт
type Account struct {
	//gorm.Model
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json: "name"`
	//Token    string `json:"token" ;sql:"-"`
}
