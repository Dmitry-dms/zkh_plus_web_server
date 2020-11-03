package models

//Аккаунт
type User struct {
	Id         int    `json:"-" db:"user_id"`
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
	FullName   string
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	CompanyId  int
}
