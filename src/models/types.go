package models



type TransactionType int


const (
	SELL TransactionType = iota+1
	BUY
)

func (t TransactionType) String() string{
	return [...]string{"SELL", "BUY"}[t-1]
}

