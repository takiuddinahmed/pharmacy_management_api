package models



type Company struct {
	Id uint `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name" gorm:"type:varchar(50);unique;not null"`
	Description string `json:"description"`
}