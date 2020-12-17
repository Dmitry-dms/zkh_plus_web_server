package models

type Notification struct {
	Id          int    `json:"-" db:"id"`
	CompanyId   int    `json:"company_id" db:"company_id"`
	Article     string `json:"article" binding:"required" db:"article"`
	Description string `json:"description" binding:"required" db:"description"`
	FullDate    string `json:"date_full" db:"date_full"`
}
