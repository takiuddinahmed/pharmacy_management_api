package models

type DrugTansaction struct{
	Id uint 
	Transaction Transaction
	Drug Drug
	Quantity int
	TransactionType TransactionType
	Price int
}