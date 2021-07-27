package models

type Customer struct {
	Id uint  `json:"id"`
	Name string `json:"name" gorm:"not null"`
	Info string `json:"info"`
	Phone string `json:"phone"`
	Address string 	`json:"address"`	
	Due float64 `json:"due"` 
}