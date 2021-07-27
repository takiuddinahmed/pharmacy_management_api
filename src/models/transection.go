package models


type Transaction struct{
	Id uint
	TransactionType TransactionType
	Customer Customer
	Price int
	Paid int
}