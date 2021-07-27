package models 


type Drug struct {
	Id uint `json:"id"`
	Name string `json:"name" gorm:"unique;not null;type:varchar(50)"`
	GenericNameId int `json:"generic_name_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GenericName GenericName `json:"generic_name,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyID int `json:"company_id"`
	Company Company `json:"company,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UnitPrice float64 `json:"unit_price"`
	Stock int `json:"stock"`
}