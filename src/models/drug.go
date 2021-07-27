package models 


type Drug struct {
	Id uint
	Name string
	GenericName GenericName
	Company company
	UnitPrice float64
	Stock int 
}