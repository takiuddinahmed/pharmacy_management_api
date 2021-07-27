package models 


type Drug struct {
	Id uint `json:"id"`
	Name string `json:"name" gorm:"unique;not null;type:varchar(50)"`
	GenericNameId int `json:"generic_name_id"`
	GenericName GenericName `json:"generic_name" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CompanyID int `json:"company_id"`
	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UnitPrice float64 `json:"unit_price"`
	Stock int `json:"stock"`
}