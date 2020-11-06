package models

type Company struct {
	Id           int    `json:"company_id" db:"company_id"`
	Name         string `json:"name" binding:"required" db:"company_name"`
	DirectorName string `json:"director_full_name" binding:"required" db:"director_full_name"`
	Phone        string `json:"company_phone" binding:"required" db:"company_phone"`
	City         string `json:"company_city" binding:"required" db:"company_city"`
	Street       string `json:"company_street" binding:"required" db:"company_street"`
	HomeNumber   string `json:"company_home_number" binding:"required" db:"company_home_number"`
	Flat         string `json:"company_flat" binding:"required" db:"company_flat"`
}
