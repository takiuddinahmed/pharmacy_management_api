package models

type GenericName struct {
	Id uint `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name" gorm:"unique;not null;type:varchar(50)"`
}